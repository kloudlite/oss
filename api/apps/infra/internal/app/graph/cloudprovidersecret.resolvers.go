package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"fmt"
	"time"

	"github.com/kloudlite/operator/pkg/operator"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"kloudlite.io/apps/infra/internal/app/graph/generated"
	"kloudlite.io/apps/infra/internal/entities"
	fn "kloudlite.io/pkg/functions"
)

// CreationTime is the resolver for the creationTime field.
func (r *cloudProviderSecretResolver) CreationTime(ctx context.Context, obj *entities.CloudProviderSecret) (string, error) {
	if obj == nil || obj.CreationTime.IsZero() {
		return "", fmt.Errorf("CloudProviderSecret object is nil")
	}
	return obj.CreationTime.Format(time.RFC3339), nil
}

// Data is the resolver for the data field.
func (r *cloudProviderSecretResolver) Data(ctx context.Context, obj *entities.CloudProviderSecret) (map[string]interface{}, error) {
	if obj == nil {
		return nil, fmt.Errorf("CloudProviderSecret object is nil")
	}
	var m map[string]any
	if err := fn.JsonConversion(obj.Data, &m); err != nil {
		return nil, err
	}
	return m, nil
}

// Enabled is the resolver for the enabled field.
func (r *cloudProviderSecretResolver) Enabled(ctx context.Context, obj *entities.CloudProviderSecret) (*bool, error) {
	panic(fmt.Errorf("not implemented: Enabled - enabled"))
}

// ID is the resolver for the id field.
func (r *cloudProviderSecretResolver) ID(ctx context.Context, obj *entities.CloudProviderSecret) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("CloudProviderSecret object is nil")
	}

	return string(obj.Id), nil
}

// Status is the resolver for the status field.
func (r *cloudProviderSecretResolver) Status(ctx context.Context, obj *entities.CloudProviderSecret) (*operator.Status, error) {
	return &operator.Status{
		IsReady: true,
	}, nil
}

// StringData is the resolver for the stringData field.
func (r *cloudProviderSecretResolver) StringData(ctx context.Context, obj *entities.CloudProviderSecret) (map[string]interface{}, error) {
	if obj == nil {

		return nil, fmt.Errorf("CloudProviderSecret object is nil")
	}

	var m map[string]any
	if err := fn.JsonConversion(obj.StringData, &m); err != nil {
		return nil, err
	}

	return m, nil
}

// Type is the resolver for the type field.
func (r *cloudProviderSecretResolver) Type(ctx context.Context, obj *entities.CloudProviderSecret) (*string, error) {
	if obj == nil {
		return nil, fmt.Errorf("CloudProviderSecret object is nil")
	}
	return fn.New(string(obj.Type)), nil
}

// UpdateTime is the resolver for the updateTime field.
func (r *cloudProviderSecretResolver) UpdateTime(ctx context.Context, obj *entities.CloudProviderSecret) (string, error) {
	if obj == nil || obj.UpdateTime.IsZero() {
		return "", fmt.Errorf("CloudProviderSecret object is nil")
	}

	return obj.UpdateTime.Format(time.RFC3339), nil
}

// Data is the resolver for the data field.
func (r *cloudProviderSecretInResolver) Data(ctx context.Context, obj *entities.CloudProviderSecret, data map[string]interface{}) error {
	if obj == nil {
		return fmt.Errorf("CloudProviderSecret object is nil")
	}

	return fn.JsonConversion(data, &obj.Data)
}

// Enabled is the resolver for the enabled field.
func (r *cloudProviderSecretInResolver) Enabled(ctx context.Context, obj *entities.CloudProviderSecret, data *bool) error {
	panic(fmt.Errorf("not implemented: Enabled - enabled"))
}

// Metadata is the resolver for the metadata field.
func (r *cloudProviderSecretInResolver) Metadata(ctx context.Context, obj *entities.CloudProviderSecret, data *v1.ObjectMeta) error {
	if obj == nil {
		return fmt.Errorf("CloudProviderSecret object is nil")
	}

	return fn.JsonConversion(data, &obj.ObjectMeta)
}

// StringData is the resolver for the stringData field.
func (r *cloudProviderSecretInResolver) StringData(ctx context.Context, obj *entities.CloudProviderSecret, data map[string]interface{}) error {
	if obj == nil {
		return fmt.Errorf("CloudProviderSecret object is nil")
	}

	return fn.JsonConversion(data, &obj.StringData)
}

// Type is the resolver for the type field.
func (r *cloudProviderSecretInResolver) Type(ctx context.Context, obj *entities.CloudProviderSecret, data *string) error {
	if obj == nil {
		return fmt.Errorf("CloudProviderSecret object is nil")
	}

	if data == nil {
		return fmt.Errorf("data is nil")
	}

	obj.Type = corev1.SecretType(*data)
	return nil
}

// CloudProviderSecret returns generated.CloudProviderSecretResolver implementation.
func (r *Resolver) CloudProviderSecret() generated.CloudProviderSecretResolver {
	return &cloudProviderSecretResolver{r}
}

// CloudProviderSecretIn returns generated.CloudProviderSecretInResolver implementation.
func (r *Resolver) CloudProviderSecretIn() generated.CloudProviderSecretInResolver {
	return &cloudProviderSecretInResolver{r}
}

type cloudProviderSecretResolver struct{ *Resolver }
type cloudProviderSecretInResolver struct{ *Resolver }
