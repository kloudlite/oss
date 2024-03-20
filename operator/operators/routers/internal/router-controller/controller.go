package router_controller

import (
	"context"
	"fmt"
	"time"

	certmanagerv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	certmanagermetav1 "github.com/cert-manager/cert-manager/pkg/apis/meta/v1"
	"golang.org/x/crypto/bcrypt"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	apiErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	crdsv1 "github.com/kloudlite/operator/apis/crds/v1"
	"github.com/kloudlite/operator/operators/routers/internal/env"
	"github.com/kloudlite/operator/operators/routers/internal/templates"
	"github.com/kloudlite/operator/pkg/constants"
	fn "github.com/kloudlite/operator/pkg/functions"
	"github.com/kloudlite/operator/pkg/kubectl"
	"github.com/kloudlite/operator/pkg/logging"
	rApi "github.com/kloudlite/operator/pkg/operator"
	stepResult "github.com/kloudlite/operator/pkg/operator/step-result"
)

type Reconciler struct {
	client.Client
	Scheme     *runtime.Scheme
	logger     logging.Logger
	Name       string
	Env        *env.Env
	yamlClient kubectl.YAMLClient

	templateIngress []byte
}

func (r *Reconciler) GetName() string {
	return r.Name
}

const (
	IngressReady    string = "ingress-ready"
	BasicAuthReady  string = "basic-auth-ready"
	DefaultsPatched string = "patch-defaults"

	Finalizing         string = "finalizing"
	CheckHttpsCerteady string = "https-cert-ready"

	EnsuringHttpsCertsIfEnabled string = "ensuring-https-certs-if-enabled"
	SettingUpBasicAuthIfEnabled string = "setting-up-basic-auth-if-enabled"
	CreatingIngressResources    string = "creating-ingress-resources"

	CleaningUpResources string = "cleaning-up-resourcess"
)

var (
	ApplyChecklist = []rApi.CheckMeta{
		{Name: DefaultsPatched, Title: "Defaults Patched"},
		{Name: EnsuringHttpsCertsIfEnabled, Title: "Ensuring HTTPS Cert if enabled"},
		{Name: SettingUpBasicAuthIfEnabled, Title: "Setting Up Basic Auth if enabled"},
	}

	DeleteChecklist = []rApi.CheckMeta{
		{Name: CleaningUpResources, Title: "Cleaning Up Resources"},
	}
)

// +kubebuilder:rbac:groups=crds.kloudlite.io,resources=crds,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=crds.kloudlite.io,resources=crds/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=crds.kloudlite.io,resources=crds/finalizers,verbs=update

func (r *Reconciler) Reconcile(ctx context.Context, request ctrl.Request) (ctrl.Result, error) {
	req, err := rApi.NewRequest(rApi.NewReconcilerCtx(ctx, r.logger), r.Client, request.NamespacedName, &crdsv1.Router{})
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if !req.ShouldReconcile() {
		return ctrl.Result{}, nil
	}

	req.PreReconcile()
	defer req.PostReconcile()

	if req.Object.GetDeletionTimestamp() != nil {
		if x := r.finalize(req); !x.ShouldProceed() {
			return x.ReconcilerResponse()
		}
		return ctrl.Result{}, nil
	}

	if step := req.ClearStatusIfAnnotated(); !step.ShouldProceed() {
		return step.ReconcilerResponse()
	}

	if step := req.EnsureCheckList(ApplyChecklist); !step.ShouldProceed() {
		return step.ReconcilerResponse()
	}

	if step := req.EnsureLabelsAndAnnotations(); !step.ShouldProceed() {
		return step.ReconcilerResponse()
	}

	if step := req.EnsureFinalizers(constants.ForegroundFinalizer, constants.CommonFinalizer); !step.ShouldProceed() {
		return step.ReconcilerResponse()
	}

	if step := r.patchDefaults(req); !step.ShouldProceed() {
		return step.ReconcilerResponse()
	}

	if step := r.EnsuringHttpsCerts(req); !step.ShouldProceed() {
		return step.ReconcilerResponse()
	}

	if step := r.reconBasicAuth(req); !step.ShouldProceed() {
		return step.ReconcilerResponse()
	}

	if step := r.ensureIngresses(req); !step.ShouldProceed() {
		return step.ReconcilerResponse()
	}

	req.Object.Status.IsReady = true
	return ctrl.Result{}, nil
}

func (r *Reconciler) patchDefaults(req *rApi.Request[*crdsv1.Router]) stepResult.Result {
	ctx, obj, checks := req.Context(), req.Object, req.Object.Status.Checks
	check := rApi.Check{Generation: obj.Generation}

	req.LogPreCheck(DefaultsPatched)
	defer req.LogPostCheck(DefaultsPatched)

	hasUpdated := false

	if obj.Spec.BasicAuth != nil && obj.Spec.BasicAuth.Enabled && obj.Spec.BasicAuth.SecretName == "" {
		hasUpdated = true
		obj.Spec.BasicAuth.SecretName = obj.Name + "-basic-auth"
	}

	if hasUpdated {
		if err := r.Update(ctx, obj); err != nil {
			return req.CheckFailed(DefaultsPatched, check, err.Error()).Err(nil)
		}
		return req.Done().RequeueAfter(1 * time.Second)
	}

	check.Status = true
	if check != checks[DefaultsPatched] {
		checks[DefaultsPatched] = check
		if sr := req.UpdateStatus(); !sr.ShouldProceed() {
			return sr
		}
	}
	return req.Next()
}

func (r *Reconciler) finalize(req *rApi.Request[*crdsv1.Router]) stepResult.Result {
	obj := req.Object
	check := rApi.Check{Generation: obj.Generation}

	checkName := CleaningUpResources

	req.LogPreCheck(checkName)
	defer req.LogPostCheck(checkName)

	fail := func(err error) stepResult.Result {
		return req.CheckFailed(checkName, check, err.Error())
	}

	if step := req.CleanupOwnedResources(); !step.ShouldProceed() {
		_, err := step.ReconcilerResponse()
		return fail(err)
	}

	return req.Finalize()
}

func genTLSCertName(domain string) string {
	return fmt.Sprintf("%s-tls", domain)
}

func (r *Reconciler) getRouterClusterIssuer(obj *crdsv1.Router) string {
	https := obj.Spec.Https

	if https != nil && https.ClusterIssuer != "" {
		return https.ClusterIssuer
	}

	return r.Env.DefaultClusterIssuer
}

func isHttpsEnabled(obj *crdsv1.Router) bool {
	return obj.Spec.Https != nil && obj.Spec.Https.Enabled
}

func (r *Reconciler) EnsuringHttpsCerts(req *rApi.Request[*crdsv1.Router]) stepResult.Result {
	ctx, obj := req.Context(), req.Object
	check := rApi.Check{Generation: obj.Generation, State: rApi.RunningState}

	checkName := EnsuringHttpsCertsIfEnabled

	req.LogPreCheck(checkName)
	defer req.LogPostCheck(checkName)

	fail := func(err error) stepResult.Result {
		return req.CheckFailed(checkName, check, err.Error())
	}

	if isHttpsEnabled(obj) {
		_, nonWildcardDomains, err := r.parseAndExtractDomains(req)
		if err != nil {
			return fail(err)
		}

		for _, domain := range nonWildcardDomains {
			cert, err := rApi.Get(ctx, r.Client, fn.NN(r.Env.CertificateNamespace, genTLSCertName(domain)), &certmanagerv1.Certificate{})
			if err != nil {
				if !apiErrors.IsNotFound(err) {
					return fail(err)
				}
				cert = nil
			}

			if cert == nil {
				cert := &certmanagerv1.Certificate{
					TypeMeta: metav1.TypeMeta{
						Kind:       "Certificate",
						APIVersion: certmanagerv1.SchemeGroupVersion.String(),
					},
					ObjectMeta: metav1.ObjectMeta{
						Name:      genTLSCertName(domain),
						Namespace: r.Env.CertificateNamespace,
					},
					Spec: certmanagerv1.CertificateSpec{
						DNSNames: []string{domain},
						IssuerRef: certmanagermetav1.ObjectReference{
							Name:  r.getRouterClusterIssuer(obj),
							Kind:  "ClusterIssuer",
							Group: certmanagerv1.SchemeGroupVersion.Group,
						},
						RenewBefore: &metav1.Duration{
							Duration: 15 * 24 * time.Hour, // 15 days prior
						},
						SecretName: genTLSCertName(domain),
						Usages: []certmanagerv1.KeyUsage{
							certmanagerv1.UsageDigitalSignature,
							certmanagerv1.UsageKeyEncipherment,
						},
					},
				}
				if err := r.Create(ctx, cert); err != nil {
					return fail(err)
				}
			}

			if _, err := IsHttpsCertReady(cert); err != nil {
				return fail(err)
			}

			certSecret, err := rApi.Get(ctx, r.Client, fn.NN(r.Env.CertificateNamespace, genTLSCertName(domain)), &corev1.Secret{})
			if err != nil {
				return fail(err)
			}

			copyTLSSecret := &corev1.Secret{ObjectMeta: metav1.ObjectMeta{Name: genTLSCertName(domain), Namespace: obj.Namespace}, Type: corev1.SecretTypeTLS}
			if _, err := controllerutil.CreateOrUpdate(ctx, r.Client, copyTLSSecret, func() error {
				if copyTLSSecret.Annotations == nil {
					copyTLSSecret.Annotations = make(map[string]string, 1)
				}
				copyTLSSecret.Annotations["kloudlite.io/secret.cloned-by"] = "router"
				// copyTLSSecret.SetOwnerReferences([]metav1.OwnerReference{fn.AsOwner(obj, true)})

				copyTLSSecret.Data = certSecret.Data
				copyTLSSecret.StringData = certSecret.StringData
				return nil
			}); err != nil {
				return fail(err)
			}
		}
	}

	check.Status = true
	if check != obj.Status.Checks[checkName] {
		fn.MapSet(&obj.Status.Checks, checkName, check)
		if sr := req.UpdateStatus(); !sr.ShouldProceed() {
			return sr
		}
	}

	return req.Next()
}

func (r *Reconciler) reconBasicAuth(req *rApi.Request[*crdsv1.Router]) stepResult.Result {
	ctx, obj := req.Context(), req.Object
	check := rApi.Check{Generation: obj.Generation, State: rApi.RunningState}

	checkName := SettingUpBasicAuthIfEnabled

	req.LogPreCheck(checkName)
	defer req.LogPostCheck(checkName)

	fail := func(err error) stepResult.Result {
		return req.CheckFailed(checkName, check, err.Error())
	}

	if obj.Spec.BasicAuth != nil && obj.Spec.BasicAuth.Enabled {
		if len(obj.Spec.BasicAuth.Username) == 0 {
			return req.CheckFailed(BasicAuthReady, check, ".spec.basicAuth.username must be defined when .spec.basicAuth.enabled is set to true").Err(nil)
		}

		basicAuthScrt := &corev1.Secret{ObjectMeta: metav1.ObjectMeta{Name: obj.Spec.BasicAuth.SecretName, Namespace: obj.Namespace}, Type: "Opaque"}

		if _, err := controllerutil.CreateOrUpdate(ctx, r.Client, basicAuthScrt, func() error {
			basicAuthScrt.SetOwnerReferences([]metav1.OwnerReference{fn.AsOwner(obj, true)})
			if _, ok := basicAuthScrt.Data["password"]; ok {
				return nil
			}

			password := fn.CleanerNanoid(48)
			ePass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				return err
			}
			basicAuthScrt.Data = map[string][]byte{
				"auth":     []byte(fmt.Sprintf("%s:%s", obj.Spec.BasicAuth.Username, ePass)),
				"username": []byte(obj.Spec.BasicAuth.Username),
				"password": []byte(password),
			}
			return nil
		}); err != nil {
			return fail(err)
		}

		req.AddToOwnedResources(rApi.ParseResourceRef(basicAuthScrt))
	}

	check.Status = true
	if check != obj.Status.Checks[checkName] {
		fn.MapSet(&obj.Status.Checks, checkName, check)
		if sr := req.UpdateStatus(); !sr.ShouldProceed() {
			return sr
		}
	}

	return req.Next()
}

func (r *Reconciler) ensureIngresses(req *rApi.Request[*crdsv1.Router]) stepResult.Result {
	ctx, obj := req.Context(), req.Object
	check := rApi.Check{Generation: obj.Generation, State: rApi.RunningState}

	checkName := CreatingIngressResources

	req.LogPreCheck(checkName)
	defer req.LogPostCheck(checkName)

	fail := func(err error) stepResult.Result {
		return req.CheckFailed(checkName, check, err.Error())
	}

	if len(obj.Spec.Routes) == 0 {
		check.Status = true
		fn.MapSet(&obj.Status.Checks, IngressReady, check)
		if err := r.Status().Update(ctx, obj); err != nil {
			return req.CheckFailed(IngressReady, check, err.Error())
		}
	}

	wcDomains, nonWcDomains, err := r.parseAndExtractDomains(req)
	if err != nil {
		return req.CheckFailed(IngressReady, check, err.Error())
	}

	nginxIngressAnnotations := GenNginxIngressAnnotations(obj)

	b, err := templates.ParseBytes(
		r.templateIngress, map[string]any{
			"name":      obj.Name,
			"namespace": obj.Namespace,

			"owner-refs":  []metav1.OwnerReference{fn.AsOwner(obj, true)},
			"labels":      obj.GetLabels(),
			"annotations": nginxIngressAnnotations,

			"non-wildcard-domains": nonWcDomains,
			"wildcard-domains":     wcDomains,
			"router-domains":       obj.Spec.Domains,

			"ingress-class": func() string {
				if obj.Spec.IngressClass != "" {
					return obj.Spec.IngressClass
				}
				return r.Env.DefaultIngressClass
			}(),
			"cluster-issuer": func() string {
				if obj.Spec.Https != nil && obj.Spec.Https.ClusterIssuer != "" {
					return obj.Spec.Https.ClusterIssuer
				}
				return r.Env.DefaultClusterIssuer
			}(),

			"routes": obj.Spec.Routes,

			"is-https-enabled": isHttpsEnabled(obj),
		},
	)
	if err != nil {
		return fail(err).Err(nil)
	}

	rr, err := r.yamlClient.ApplyYAML(ctx, b)
	if err != nil {
		return fail(err)
	}

	req.AddToOwnedResources(rr...)

	check.Status = true

	if check != obj.Status.Checks[checkName] {
		fn.MapSet(&obj.Status.Checks, checkName, check)
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
	r.yamlClient = kubectl.NewYAMLClientOrDie(mgr.GetConfig(), kubectl.YAMLClientOpts{Logger: r.logger})

	var err error
	r.templateIngress, err = templates.ReadIngressTemplate()
	if err != nil {
		return err
	}

	builder := ctrl.NewControllerManagedBy(mgr).For(&crdsv1.Router{})
	builder.Owns(&networkingv1.Ingress{})
	builder.Owns(&certmanagerv1.Certificate{})

	// builder.Watches(&certmanagerv1.Certificate{},
	// 	handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, obj client.Object) []reconcile.Request {
	// 		if obj.GetNamespace() != r.Env.CertificateNamespace {
	// 			return nil
	// 		}
	//
	// 		rlist, err := kubectl.PaginatedList[*crdsv1.Router](ctx, r.Client, &crdsv1.RouterList{}, &client.ListOptions{Limit: 10})
	// 		if err != nil {
	// 			return nil
	// 		}
	//
	// 		reqs := make([]reconcile.Request, 0, 10)
	//
	// 		for router := range rlist {
	// 			reqs = append(reqs, reconcile.Request{NamespacedName: fn.NN(router.Namespace, router.Name)})
	// 		}
	//
	// 		return reqs
	// 	}))

	builder.WithOptions(controller.Options{MaxConcurrentReconciles: r.Env.MaxConcurrentReconciles})
	builder.WithEventFilter(rApi.ReconcileFilter())
	return builder.Complete(r)
}
