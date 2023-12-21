package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"github.com/kloudlite/api/pkg/errors"
	"time"

	"github.com/kloudlite/api/apps/infra/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/infra/internal/app/graph/model"
	"github.com/kloudlite/api/common"
	fn "github.com/kloudlite/api/pkg/functions"
	"github.com/kloudlite/api/pkg/types"
	"github.com/kloudlite/operator/pkg/operator"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Checks is the resolver for the checks field.
func (r *github__com___kloudlite___operator___pkg___operator__StatusResolver) Checks(ctx context.Context, obj *operator.Status) (map[string]interface{}, error) {
	var m map[string]any
	if err := fn.JsonConversion(obj.Checks, &m); err != nil {
		return nil, errors.NewE(err)
	}
	return m, nil
}

// LastReconcileTime is the resolver for the lastReconcileTime field.
func (r *github__com___kloudlite___operator___pkg___operator__StatusResolver) LastReconcileTime(ctx context.Context, obj *operator.Status) (*string, error) {
	if obj == nil {
		return nil, errors.Newf("syncStatus is nil")
	}
	if obj.LastReconcileTime == nil {
		return fn.New(time.Now().Format(time.RFC3339)), nil
	}
	return fn.New(obj.LastReconcileTime.Format(time.RFC3339)), nil
}

// Message is the resolver for the message field.
func (r *github__com___kloudlite___operator___pkg___operator__StatusResolver) Message(ctx context.Context, obj *operator.Status) (*model.GithubComKloudliteOperatorPkgRawJSONRawJSON, error) {
	if obj == nil {
		return nil, errors.Newf("syncStatus is nil")
	}
	if obj.Message == nil {
		return nil, nil
	}
	return &model.GithubComKloudliteOperatorPkgRawJSONRawJSON{
		RawMessage: obj.Message.RawMessage,
	}, nil
}

// UserID is the resolver for the userId field.
func (r *kloudlite__io___common__CreatedOrUpdatedByResolver) UserID(ctx context.Context, obj *common.CreatedOrUpdatedBy) (string, error) {
	if obj == nil {
		return "", errors.Newf("createdOrUpdatedBy is nil")
	}
	return string(obj.UserId), nil
}

// LastSyncedAt is the resolver for the lastSyncedAt field.
func (r *kloudlite__io___pkg___types__SyncStatusResolver) LastSyncedAt(ctx context.Context, obj *types.SyncStatus) (*string, error) {
	if obj == nil {
		return nil, errors.Newf("syncStatus is nil")
	}
	return fn.New(obj.LastSyncedAt.Format(time.RFC3339)), nil
}

// SyncScheduledAt is the resolver for the syncScheduledAt field.
func (r *kloudlite__io___pkg___types__SyncStatusResolver) SyncScheduledAt(ctx context.Context, obj *types.SyncStatus) (*string, error) {
	if obj == nil {
		return nil, errors.Newf("syncStatus is nil")
	}
	return fn.New(obj.SyncScheduledAt.Format(time.RFC3339)), nil
}

// Annotations is the resolver for the annotations field.
func (r *metadataResolver) Annotations(ctx context.Context, obj *v1.ObjectMeta) (map[string]interface{}, error) {
	var m map[string]any
	if err := fn.JsonConversion(obj.Annotations, &m); err != nil {
		return nil, errors.NewE(err)
	}
	return m, nil
}

// CreationTimestamp is the resolver for the creationTimestamp field.
func (r *metadataResolver) CreationTimestamp(ctx context.Context, obj *v1.ObjectMeta) (string, error) {
	if obj == nil {
		return "", errors.Newf("metadata is nil")
	}
	return obj.CreationTimestamp.Format(time.RFC3339), nil
}

// DeletionTimestamp is the resolver for the deletionTimestamp field.
func (r *metadataResolver) DeletionTimestamp(ctx context.Context, obj *v1.ObjectMeta) (*string, error) {
	if obj == nil {
		return nil, errors.Newf("metadata is nil")
	}

	if obj.DeletionTimestamp == nil {
		return nil, nil
	}

	return fn.New(obj.DeletionTimestamp.Format(time.RFC3339)), nil
}

// Labels is the resolver for the labels field.
func (r *metadataResolver) Labels(ctx context.Context, obj *v1.ObjectMeta) (map[string]interface{}, error) {
	var m map[string]any
	if err := fn.JsonConversion(obj.Labels, &m); err != nil {
		return nil, errors.NewE(err)
	}
	return m, nil
}

// Annotations is the resolver for the annotations field.
func (r *metadataInResolver) Annotations(ctx context.Context, obj *v1.ObjectMeta, data map[string]interface{}) error {
	var m map[string]string
	if err := fn.JsonConversion(data, &m); err != nil {
		return errors.NewE(err)
	}
	obj.SetAnnotations(m)
	return nil
}

// Labels is the resolver for the labels field.
func (r *metadataInResolver) Labels(ctx context.Context, obj *v1.ObjectMeta, data map[string]interface{}) error {
	var m map[string]string
	if err := fn.JsonConversion(data, &m); err != nil {
		return errors.NewE(err)
	}
	obj.SetLabels(m)
	return nil
}

// Github__com___kloudlite___operator___pkg___operator__Status returns generated.Github__com___kloudlite___operator___pkg___operator__StatusResolver implementation.
func (r *Resolver) Github__com___kloudlite___operator___pkg___operator__Status() generated.Github__com___kloudlite___operator___pkg___operator__StatusResolver {
	return &github__com___kloudlite___operator___pkg___operator__StatusResolver{r}
}

// Kloudlite__io___common__CreatedOrUpdatedBy returns generated.Kloudlite__io___common__CreatedOrUpdatedByResolver implementation.
func (r *Resolver) Kloudlite__io___common__CreatedOrUpdatedBy() generated.Kloudlite__io___common__CreatedOrUpdatedByResolver {
	return &kloudlite__io___common__CreatedOrUpdatedByResolver{r}
}

// Kloudlite__io___pkg___types__SyncStatus returns generated.Kloudlite__io___pkg___types__SyncStatusResolver implementation.
func (r *Resolver) Kloudlite__io___pkg___types__SyncStatus() generated.Kloudlite__io___pkg___types__SyncStatusResolver {
	return &kloudlite__io___pkg___types__SyncStatusResolver{r}
}

// Metadata returns generated.MetadataResolver implementation.
func (r *Resolver) Metadata() generated.MetadataResolver { return &metadataResolver{r} }

// MetadataIn returns generated.MetadataInResolver implementation.
func (r *Resolver) MetadataIn() generated.MetadataInResolver { return &metadataInResolver{r} }

type github__com___kloudlite___operator___pkg___operator__StatusResolver struct{ *Resolver }
type kloudlite__io___common__CreatedOrUpdatedByResolver struct{ *Resolver }
type kloudlite__io___pkg___types__SyncStatusResolver struct{ *Resolver }
type metadataResolver struct{ *Resolver }
type metadataInResolver struct{ *Resolver }
