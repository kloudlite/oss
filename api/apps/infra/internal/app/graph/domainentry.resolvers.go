package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"fmt"
	"time"

	"github.com/kloudlite/api/apps/infra/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/infra/internal/entities"
)

// CreationTime is the resolver for the creationTime field.
func (r *domainEntryResolver) CreationTime(ctx context.Context, obj *entities.DomainEntry) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("domainEntry is nil")
	}
	return obj.CreationTime.Format(time.RFC3339), nil
}

// ID is the resolver for the id field.
func (r *domainEntryResolver) ID(ctx context.Context, obj *entities.DomainEntry) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("domainEntry is nil")
	}
	return string(obj.Id), nil
}

// UpdateTime is the resolver for the updateTime field.
func (r *domainEntryResolver) UpdateTime(ctx context.Context, obj *entities.DomainEntry) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("domainEntry is nil")
	}

	return obj.UpdateTime.Format(time.RFC3339), nil
}

// DomainEntry returns generated.DomainEntryResolver implementation.
func (r *Resolver) DomainEntry() generated.DomainEntryResolver { return &domainEntryResolver{r} }

type domainEntryResolver struct{ *Resolver }
