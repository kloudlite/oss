package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"fmt"
	"time"

	"kloudlite.io/apps/container-registry/internal/app/graph/generated"
	"kloudlite.io/apps/container-registry/internal/app/graph/model"
	"kloudlite.io/apps/container-registry/internal/domain/entities"
)

// CreationTime is the resolver for the creationTime field.
func (r *tagResolver) CreationTime(ctx context.Context, obj *entities.Tag) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("resource is nil")
	}
	return obj.CreationTime.Format(time.RFC3339), nil
}

// ID is the resolver for the id field.
func (r *tagResolver) ID(ctx context.Context, obj *entities.Tag) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("resource is nil")
	}
	return string(obj.Id), nil
}

// References is the resolver for the references field.
func (r *tagResolver) References(ctx context.Context, obj *entities.Tag) ([]*model.KloudliteIoAppsContainerRegistryInternalDomainEntitiesRepoReference, error) {
	if obj == nil {
		return nil, fmt.Errorf("resource is nil")
	}

	references := make([]*model.KloudliteIoAppsContainerRegistryInternalDomainEntitiesRepoReference, len(obj.References))
	for i, reference := range obj.References {
		references[i] = &model.KloudliteIoAppsContainerRegistryInternalDomainEntitiesRepoReference{
			Digest:    reference.Digest,
			MediaType: reference.MediaType,
			Size:      reference.Size,
		}
	}
	return references, nil
}

// UpdateTime is the resolver for the updateTime field.
func (r *tagResolver) UpdateTime(ctx context.Context, obj *entities.Tag) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("resource is nil")
	}
	return obj.UpdateTime.Format(time.RFC3339), nil
}

// Tag returns generated.TagResolver implementation.
func (r *Resolver) Tag() generated.TagResolver { return &tagResolver{r} }

type tagResolver struct{ *Resolver }
