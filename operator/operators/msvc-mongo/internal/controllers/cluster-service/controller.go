package clusterService

import (
	"context"
	"encoding/json"
	"time"

	"github.com/kloudlite/operator/operators/msvc-mongo/internal/env"

	mongodbMsvcv1 "github.com/kloudlite/operator/apis/mongodb.msvc/v1"
	"github.com/kloudlite/operator/pkg/conditions"
	"github.com/kloudlite/operator/pkg/constants"
	fn "github.com/kloudlite/operator/pkg/functions"
	"github.com/kloudlite/operator/pkg/harbor"
	"github.com/kloudlite/operator/pkg/kubectl"
	"github.com/kloudlite/operator/pkg/logging"
	rApi "github.com/kloudlite/operator/pkg/operator"
	stepResult "github.com/kloudlite/operator/pkg/operator/step-result"
	"github.com/kloudlite/operator/pkg/templates"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

type Reconciler struct {
	client.Client
	Scheme     *runtime.Scheme
	env        *env.Env
	harborCli  *harbor.Client
	logger     logging.Logger
	Name       string
	yamlClient kubectl.YAMLClient
}

func (r *Reconciler) GetName() string {
	return r.Name
}

const (
	HelmReady        string = "helm-ready"
	StsReady         string = "sts-ready"
	AccessCredsReady string = "access-creds-ready"
)

const (
	KeyRootPassword string = "root-password"
)

// +kubebuilder:rbac:groups=mongodb.msvc.kloudlite.io,resources=clusterServices,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=mongodb.msvc.kloudlite.io,resources=clusterServices/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=mongodb.msvc.kloudlite.io,resources=clusterServices/finalizers,verbs=update

func (r *Reconciler) Reconcile(ctx context.Context, request ctrl.Request) (ctrl.Result, error) {
	req, err := rApi.NewRequest(context.WithValue(ctx, "logger", r.logger), r.Client, request.NamespacedName, &mongodbMsvcv1.ClusterService{})
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if req.Object.GetDeletionTimestamp() != nil {
		if x := r.finalize(req); !x.ShouldProceed() {
			return x.ReconcilerResponse()
		}
		return ctrl.Result{}, nil
	}

	req.LogPreReconcile()
	defer req.LogPostReconcile()

	if step := req.ClearStatusIfAnnotated(); !step.ShouldProceed() {
		return step.ReconcilerResponse()
	}

	if step := req.RestartIfAnnotated(); !step.ShouldProceed() {
		return step.ReconcilerResponse()
	}

	if step := req.EnsureChecks(HelmReady, StsReady, AccessCredsReady); !step.ShouldProceed() {
		return step.ReconcilerResponse()
	}

	if step := req.EnsureLabelsAndAnnotations(); !step.ShouldProceed() {
		return step.ReconcilerResponse()
	}

	if step := req.EnsureFinalizers(constants.ForegroundFinalizer, constants.CommonFinalizer); !step.ShouldProceed() {
		return step.ReconcilerResponse()
	}

	if step := r.reconAccessCreds(req); !step.ShouldProceed() {
		return step.ReconcilerResponse()
	}

	if step := r.reconHelm(req); !step.ShouldProceed() {
		return step.ReconcilerResponse()
	}

	if step := r.reconSts(req); !step.ShouldProceed() {
		return step.ReconcilerResponse()
	}

	req.Object.Status.IsReady = true
	return ctrl.Result{RequeueAfter: r.env.ReconcilePeriod}, r.Status().Update(ctx, req.Object)
}

func (r *Reconciler) finalize(req *rApi.Request[*mongodbMsvcv1.ClusterService]) stepResult.Result {
	return req.Finalize()
}

func (r *Reconciler) reconAccessCreds(req *rApi.Request[*mongodbMsvcv1.ClusterService]) stepResult.Result {
	ctx, obj, checks := req.Context(), req.Object, req.Object.Status.Checks
	check := rApi.Check{Generation: obj.Generation}

	req.LogPreCheck(AccessCredsReady)
	defer req.LogPostCheck(AccessCredsReady)

	secretName := "msvc-" + obj.Name
	scrt, err := rApi.Get(ctx, r.Client, fn.NN(obj.Namespace, secretName), &corev1.Secret{})
	if err != nil {
		req.Logger.Infof("secret %s does not exist yet, would be creating it ...", fn.NN(obj.Namespace, secretName).String())
	}

	if scrt == nil {
		rootPassword := fn.CleanerNanoid(40)
		b, err := templates.Parse(
			templates.Secret, map[string]any{
				"name":       secretName,
				"namespace":  obj.Namespace,
				"labels":     obj.GetLabels(),
				"owner-refs": obj.GetOwnerReferences(),
				"string-data": map[string]string{
					"ROOT_PASSWORD": rootPassword,
					// TODO: user
				},
			},
		)

		if err != nil {
			return req.CheckFailed(AccessCredsReady, check, err.Error())
		}

		if _, err := r.yamlClient.ApplyYAML(ctx, b); err != nil {
			return req.CheckFailed(AccessCredsReady, check, err.Error())
		}

		if sr := req.UpdateStatus(); !sr.ShouldProceed() {
			return sr
		}
	}

	if !fn.IsOwner(obj, fn.AsOwner(scrt)) {
		obj.SetOwnerReferences(append(obj.GetOwnerReferences(), fn.AsOwner(scrt)))
		if err := r.Update(ctx, obj); err != nil {
			return req.Done().Err(err)
		}
		return req.Done().RequeueAfter(2 * time.Second)
	}

	check.Status = true
	if check != checks[AccessCredsReady] {
		checks[AccessCredsReady] = check
		if sr := req.UpdateStatus(); !sr.ShouldProceed() {
			return sr
		}
	}

	rApi.SetLocal(req, KeyRootPassword, string(scrt.Data["ROOT_PASSWORD"]))
	return req.Next()
}

func (r *Reconciler) reconHelm(req *rApi.Request[*mongodbMsvcv1.ClusterService]) stepResult.Result {
	ctx, obj, checks := req.Context(), req.Object, req.Object.Status.Checks
	check := rApi.Check{Generation: obj.Generation}

	req.LogPreCheck(HelmReady)
	defer req.LogPostCheck(HelmReady)

	helmRes, err := rApi.Get(
		ctx, r.Client, fn.NN(obj.Namespace, obj.Name), fn.NewUnstructured(constants.HelmMongoDBType),
	)
	if err != nil {
		req.Logger.Infof("helm resource (%s) not found, will be creating it", fn.NN(obj.Namespace, obj.Name).String())
	}

	// rootPassword, ok := rApi.GetLocal[string](req, KeyRootPassword)
	// if !ok {
	// 	return req.CheckFailed(HelmReady, check, err.Error())
	// }

	if helmRes == nil || check.Generation > checks[HelmReady].Generation {
		b, err := templates.Parse(templates.MongoDBCluster, map[string]any{})

		if err != nil {
			return req.CheckFailed(HelmReady, check, err.Error()).Err(nil)
		}

		if _, err := r.yamlClient.ApplyYAML(ctx, b); err != nil {
			return req.CheckFailed(HelmReady, check, err.Error()).Err(nil)
		}

		checks[HelmReady] = check
		if sr := req.UpdateStatus(); !sr.ShouldProceed() {
			return sr
		}
	}

	cds, err := conditions.FromObject(helmRes)
	if err != nil {
		return req.CheckFailed(HelmReady, check, err.Error())
	}

	deployedC := meta.FindStatusCondition(cds, "Deployed")
	if deployedC == nil {
		return req.Done()
	}

	if deployedC.Status == metav1.ConditionFalse {
		return req.CheckFailed(HelmReady, check, deployedC.Message)
	}

	if deployedC.Status == metav1.ConditionTrue {
		check.Status = true
	}

	if check != checks[HelmReady] {
		checks[HelmReady] = check
		if sr := req.UpdateStatus(); !sr.ShouldProceed() {
			return sr
		}
	}

	return req.Next()
}

func (r *Reconciler) reconSts(req *rApi.Request[*mongodbMsvcv1.ClusterService]) stepResult.Result {
	ctx, obj, checks := req.Context(), req.Object, req.Object.Status.Checks
	check := rApi.Check{Generation: obj.Generation}

	req.LogPreCheck(StsReady)
	defer req.LogPostCheck(StsReady)

	var stsList appsv1.StatefulSetList

	if err := r.List(
		ctx, &stsList, &client.ListOptions{
			LabelSelector: labels.SelectorFromValidatedSet(map[string]string{constants.MsvcNameKey: obj.Name}),
			Namespace:     obj.Namespace,
		},
	); err != nil {
		return req.CheckFailed(StsReady, check, err.Error())
	}

	for i := range stsList.Items {
		item := stsList.Items[i]
		if item.Status.AvailableReplicas != item.Status.Replicas {
			check.Status = false

			var podsList corev1.PodList
			if err := r.List(
				ctx, &podsList, &client.ListOptions{
					LabelSelector: labels.SelectorFromValidatedSet(
						map[string]string{
							constants.MsvcNameKey: obj.Name,
						},
					),
				},
			); err != nil {
				return req.CheckFailed(StsReady, check, err.Error())
			}

			messages := rApi.GetMessagesFromPods(podsList.Items...)
			if len(messages) > 0 {
				b, err := json.Marshal(messages)
				if err != nil {
					return req.CheckFailed(StsReady, check, err.Error())
				}
				return req.CheckFailed(StsReady, check, string(b))
			}
		}
	}

	check.Status = true
	if check != checks[StsReady] {
		checks[StsReady] = check
		if sr := req.UpdateStatus(); !sr.ShouldProceed() {
			return sr
		}
	}

	return req.Next()
}

func (r *Reconciler) SetupWithManager(mgr ctrl.Manager, logger logging.Logger) error {
	r.Client = mgr.GetClient()
	r.Scheme = mgr.GetScheme()
	r.logger = logger.WithName(r.Name)
	r.yamlClient = kubectl.NewYAMLClientOrDie(mgr.GetConfig())

	builder := ctrl.NewControllerManagedBy(mgr).For(&mongodbMsvcv1.ClusterService{})
	builder.Owns(&corev1.Secret{})
	builder.Owns(fn.NewUnstructured(constants.HelmMongoDBType))

	builder.Watches(
		&source.Kind{Type: &appsv1.StatefulSet{}}, handler.EnqueueRequestsFromMapFunc(
			func(obj client.Object) []reconcile.Request {
				v, ok := obj.GetLabels()[constants.MsvcNameKey]
				if !ok {
					return nil
				}
				return []reconcile.Request{{NamespacedName: fn.NN(obj.GetNamespace(), v)}}
			},
		),
	)

	return builder.Complete(r)
}
