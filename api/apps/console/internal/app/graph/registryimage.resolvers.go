package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"time"

	"github.com/kloudlite/api/apps/console/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/console/internal/entities"
)

// CreationTime is the resolver for the creationTime field.
func (r *registryImageResolver) CreationTime(ctx context.Context, obj *entities.RegistryImage) (string, error) {
	if obj == nil {
		return "", errNilRegistryImage
	}
	return obj.BaseEntity.CreationTime.Format(time.RFC3339), nil
}

// UpdateTime is the resolver for the updateTime field.
func (r *registryImageResolver) UpdateTime(ctx context.Context, obj *entities.RegistryImage) (string, error) {
	if obj == nil {
		return "", errNilRegistryImage
	}
	return obj.BaseEntity.UpdateTime.Format(time.RFC3339), nil
}

// RegistryImage returns generated.RegistryImageResolver implementation.
func (r *Resolver) RegistryImage() generated.RegistryImageResolver { return &registryImageResolver{r} }

type registryImageResolver struct{ *Resolver }
