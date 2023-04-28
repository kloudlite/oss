package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"kloudlite.io/apps/infra/internal/app/graph/generated"
	"kloudlite.io/common"
)

func (r *Resolver) Metadata() generated.MetadataResolver     { return &common.MetadataResolver{} }
func (r *Resolver) Status() generated.StatusResolver         { return &common.StatusResolver{} }
func (r *Resolver) SyncStatus() generated.SyncStatusResolver   { return &common.SyncStatusResolver{} }
func (r *Resolver) MetadataIn() generated.MetadataInResolver { return &common.MetadataInResolver{} }
func (r *Resolver) Patch() generated.PatchResolver      { return &common.PatchResolver{} }
func (r *Resolver) PatchIn() generated.PatchInResolver    { return &common.PatchInResolver{} }
