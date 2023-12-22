package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"github.com/kloudlite/api/apps/console/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/console/internal/app/graph/model"
	"github.com/kloudlite/api/apps/console/internal/domain"
	"github.com/kloudlite/api/apps/console/internal/entities"
	"github.com/kloudlite/api/pkg/errors"
	fn "github.com/kloudlite/api/pkg/functions"
	"github.com/kloudlite/api/pkg/repos"
)

// CoreCreateProject is the resolver for the core_createProject field.
func (r *mutationResolver) CoreCreateProject(ctx context.Context, project entities.Project) (*entities.Project, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateProject(cc, project)
}

// CoreUpdateProject is the resolver for the core_updateProject field.
func (r *mutationResolver) CoreUpdateProject(ctx context.Context, project entities.Project) (*entities.Project, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateProject(cc, project)
}

// CoreDeleteProject is the resolver for the core_deleteProject field.
func (r *mutationResolver) CoreDeleteProject(ctx context.Context, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteProject(cc, name); err != nil {
		return false, nil
	}
	return true, nil
}

// CoreCreateImagePullSecret is the resolver for the core_createImagePullSecret field.
func (r *mutationResolver) CoreCreateImagePullSecret(ctx context.Context, imagePullSecretIn entities.ImagePullSecret) (*entities.ImagePullSecret, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateImagePullSecret(cc, imagePullSecretIn)
}

// CoreDeleteImagePullSecret is the resolver for the core_deleteImagePullSecret field.
func (r *mutationResolver) CoreDeleteImagePullSecret(ctx context.Context, namespace string, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteImagePullSecret(cc, namespace, name); err != nil {
		return false, nil
	}
	return true, nil
}

// CoreCreateEnvironment is the resolver for the core_createEnvironment field.
func (r *mutationResolver) CoreCreateEnvironment(ctx context.Context, env entities.Workspace) (*entities.Workspace, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateEnvironment(cc, env)
}

// CoreUpdateEnvironment is the resolver for the core_updateEnvironment field.
func (r *mutationResolver) CoreUpdateEnvironment(ctx context.Context, env entities.Workspace) (*entities.Workspace, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateEnvironment(cc, env)
}

// CoreDeleteEnvironment is the resolver for the core_deleteEnvironment field.
func (r *mutationResolver) CoreDeleteEnvironment(ctx context.Context, namespace string, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteEnvironment(cc, namespace, name); err != nil {
		return false, nil
	}
	return true, nil
}

// CoreCreateWorkspace is the resolver for the core_createWorkspace field.
func (r *mutationResolver) CoreCreateWorkspace(ctx context.Context, env entities.Workspace) (*entities.Workspace, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateWorkspace(cc, env)
}

// CoreUpdateWorkspace is the resolver for the core_updateWorkspace field.
func (r *mutationResolver) CoreUpdateWorkspace(ctx context.Context, env entities.Workspace) (*entities.Workspace, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateWorkspace(cc, env)
}

// CoreDeleteWorkspace is the resolver for the core_deleteWorkspace field.
func (r *mutationResolver) CoreDeleteWorkspace(ctx context.Context, namespace string, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteWorkspace(cc, namespace, name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreCreateApp is the resolver for the core_createApp field.
func (r *mutationResolver) CoreCreateApp(ctx context.Context, app entities.App) (*entities.App, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateApp(cc, app)
}

// CoreUpdateApp is the resolver for the core_updateApp field.
func (r *mutationResolver) CoreUpdateApp(ctx context.Context, app entities.App) (*entities.App, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateApp(cc, app)
}

// CoreDeleteApp is the resolver for the core_deleteApp field.
func (r *mutationResolver) CoreDeleteApp(ctx context.Context, namespace string, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteApp(cc, namespace, name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreCreateConfig is the resolver for the core_createConfig field.
func (r *mutationResolver) CoreCreateConfig(ctx context.Context, config entities.Config) (*entities.Config, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateConfig(cc, config)
}

// CoreUpdateConfig is the resolver for the core_updateConfig field.
func (r *mutationResolver) CoreUpdateConfig(ctx context.Context, config entities.Config) (*entities.Config, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateConfig(cc, config)
}

// CoreDeleteConfig is the resolver for the core_deleteConfig field.
func (r *mutationResolver) CoreDeleteConfig(ctx context.Context, namespace string, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteConfig(cc, namespace, name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreCreateSecret is the resolver for the core_createSecret field.
func (r *mutationResolver) CoreCreateSecret(ctx context.Context, secret entities.Secret) (*entities.Secret, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateSecret(cc, secret)
}

// CoreUpdateSecret is the resolver for the core_updateSecret field.
func (r *mutationResolver) CoreUpdateSecret(ctx context.Context, secret entities.Secret) (*entities.Secret, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateSecret(cc, secret)
}

// CoreDeleteSecret is the resolver for the core_deleteSecret field.
func (r *mutationResolver) CoreDeleteSecret(ctx context.Context, namespace string, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteSecret(cc, namespace, name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreCreateRouter is the resolver for the core_createRouter field.
func (r *mutationResolver) CoreCreateRouter(ctx context.Context, router entities.Router) (*entities.Router, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateRouter(cc, router)
}

// CoreUpdateRouter is the resolver for the core_updateRouter field.
func (r *mutationResolver) CoreUpdateRouter(ctx context.Context, router entities.Router) (*entities.Router, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateRouter(cc, router)
}

// CoreDeleteRouter is the resolver for the core_deleteRouter field.
func (r *mutationResolver) CoreDeleteRouter(ctx context.Context, namespace string, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteRouter(cc, namespace, name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreCreateManagedService is the resolver for the core_createManagedService field.
func (r *mutationResolver) CoreCreateManagedService(ctx context.Context, msvc entities.ManagedService) (*entities.ManagedService, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateManagedService(cc, msvc)
}

// CoreUpdateManagedService is the resolver for the core_updateManagedService field.
func (r *mutationResolver) CoreUpdateManagedService(ctx context.Context, msvc entities.ManagedService) (*entities.ManagedService, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateManagedService(cc, msvc)
}

// CoreDeleteManagedService is the resolver for the core_deleteManagedService field.
func (r *mutationResolver) CoreDeleteManagedService(ctx context.Context, namespace string, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteManagedService(cc, namespace, name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreCreateManagedResource is the resolver for the core_createManagedResource field.
func (r *mutationResolver) CoreCreateManagedResource(ctx context.Context, mres entities.ManagedResource) (*entities.ManagedResource, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.CreateManagedResource(cc, mres)
}

// CoreUpdateManagedResource is the resolver for the core_updateManagedResource field.
func (r *mutationResolver) CoreUpdateManagedResource(ctx context.Context, mres entities.ManagedResource) (*entities.ManagedResource, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.UpdateManagedResource(cc, mres)
}

// CoreDeleteManagedResource is the resolver for the core_deleteManagedResource field.
func (r *mutationResolver) CoreDeleteManagedResource(ctx context.Context, namespace string, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.DeleteManagedResource(cc, namespace, name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreCheckNameAvailability is the resolver for the core_checkNameAvailability field.
func (r *queryResolver) CoreCheckNameAvailability(ctx context.Context, resType domain.ResType, namespace *string, name string) (*domain.CheckNameAvailabilityOutput, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		if cc.AccountName == "" {
			return nil, errors.NewE(err)
		}
	}
	return r.Domain.CheckNameAvailability(ctx, resType, cc.AccountName, namespace, name)
}

// CoreListProjects is the resolver for the core_listProjects field.
func (r *queryResolver) CoreListProjects(ctx context.Context, clusterName *string, search *model.SearchProjects, pq *repos.CursorPagination) (*model.ProjectPaginatedRecords, error) {
	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}
		if search.MarkedForDeletion != nil {
			filter["markedForDeletion"] = *search.MarkedForDeletion
		}
	}

	cc, err := toConsoleContext(ctx)
	if err != nil {
		if cc.UserId == "" || cc.AccountName == "" {
			return nil, errors.NewE(err)
		}
	}

	p, err := r.Domain.ListProjects(ctx, cc.UserId, cc.AccountName, clusterName, filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	pe := make([]*model.ProjectEdge, len(p.Edges))
	for i := range p.Edges {
		pe[i] = &model.ProjectEdge{
			Node:   p.Edges[i].Node,
			Cursor: p.Edges[i].Cursor,
		}
	}

	m := model.ProjectPaginatedRecords{
		Edges: pe,
		PageInfo: &model.PageInfo{
			EndCursor:       &p.PageInfo.EndCursor,
			HasNextPage:     p.PageInfo.HasNextPage,
			HasPreviousPage: p.PageInfo.HasPrevPage,
			StartCursor:     &p.PageInfo.StartCursor,
		},
		TotalCount: int(p.TotalCount),
	}

	return &m, nil
}

// CoreGetProject is the resolver for the core_getProject field.
func (r *queryResolver) CoreGetProject(ctx context.Context, name string) (*entities.Project, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.GetProject(cc, name)
}

// CoreResyncProject is the resolver for the core_resyncProject field.
func (r *queryResolver) CoreResyncProject(ctx context.Context, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.ResyncProject(cc, name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreListImagePullSecrets is the resolver for the infra_listImagePullSecrets field.
func (r *queryResolver) CoreListImagePullSecrets(ctx context.Context, project model.ProjectID, scope *model.WorkspaceOrEnvID, search *model.SearchImagePullSecrets, pq *repos.CursorPagination) (*model.ImagePullSecretPaginatedRecords, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}
		if search.MarkedForDeletion != nil {
			filter["markedForDeletion"] = *search.MarkedForDeletion
		}
	}

	namespace, err := func() (string, error) {
		if scope == nil {
			return r.getNamespaceFromProjectID(ctx, project)
		}
		return r.getNamespaceFromProjectAndScope(ctx, project, *scope)
	}()
	if err != nil {
		return nil, errors.NewE(err)
	}

	pr, err := r.Domain.ListImagePullSecrets(cc, namespace, filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	ipsEdges := make([]*model.ImagePullSecretEdge, len(pr.Edges))
	for i := range pr.Edges {
		ipsEdges[i] = &model.ImagePullSecretEdge{
			Node:   pr.Edges[i].Node,
			Cursor: pr.Edges[i].Cursor,
		}
	}

	m := model.ImagePullSecretPaginatedRecords{
		Edges: ipsEdges,
		PageInfo: &model.PageInfo{
			EndCursor:       &pr.PageInfo.EndCursor,
			HasNextPage:     pr.PageInfo.HasNextPage,
			HasPreviousPage: pr.PageInfo.HasPrevPage,
			StartCursor:     &pr.PageInfo.StartCursor,
		},
		TotalCount: int(pr.TotalCount),
	}

	return &m, nil
}

// InfraGetImagePullSecret is the resolver for the infra_getImagePullSecret field.
func (r *queryResolver) CoreGetImagePullSecret(ctx context.Context, project model.ProjectID, scope *model.WorkspaceOrEnvID, name string) (*entities.ImagePullSecret, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	namespace, err := func() (string, error) {
		if scope == nil {
			return r.getNamespaceFromProjectID(ctx, project)
		}
		return r.getNamespaceFromProjectAndScope(ctx, project, *scope)
	}()
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.GetImagePullSecret(cc, namespace, name)
}

// CoreResyncImagePullSecret is the resolver for the core_resyncImagePullSecret field.
func (r *queryResolver) CoreResyncImagePullSecret(ctx context.Context, project model.ProjectID, scope *model.WorkspaceOrEnvID, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	namespace, err := func() (string, error) {
		if scope == nil {
			return r.getNamespaceFromProjectID(ctx, project)
		}
		return r.getNamespaceFromProjectAndScope(ctx, project, *scope)
	}()
	if err != nil {
		return false, errors.NewE(err)
	}

	if err := r.Domain.ResyncImagePullSecret(cc, namespace, name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreListWorkspaces is the resolver for the core_listWorkspaces field.
func (r *queryResolver) CoreListWorkspaces(ctx context.Context, project model.ProjectID, search *model.SearchWorkspaces, pq *repos.CursorPagination) (*model.WorkspacePaginatedRecords, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
		if search.ProjectName != nil {
			filter["spec.projectName"] = *search.ProjectName
		}
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}
		if search.MarkedForDeletion != nil {
			filter["markedForDeletion"] = *search.MarkedForDeletion
		}
	}

	namespace, err := r.getNamespaceFromProjectID(ctx, project)
	if err != nil {
		return nil, errors.NewE(err)
	}

	pw, err := r.Domain.ListWorkspaces(cc, namespace, filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	we := make([]*model.WorkspaceEdge, len(pw.Edges))
	for i := range pw.Edges {
		we[i] = &model.WorkspaceEdge{
			Node:   pw.Edges[i].Node,
			Cursor: pw.Edges[i].Cursor,
		}
	}

	m := model.WorkspacePaginatedRecords{
		Edges: we,
		PageInfo: &model.PageInfo{
			EndCursor:       &pw.PageInfo.EndCursor,
			HasNextPage:     pw.PageInfo.HasNextPage,
			HasPreviousPage: pw.PageInfo.HasPrevPage,
			StartCursor:     &pw.PageInfo.StartCursor,
		},
		TotalCount: int(pw.TotalCount),
	}

	return &m, nil
}

// CoreGetWorkspace is the resolver for the core_getWorkspace field.
func (r *queryResolver) CoreGetWorkspace(ctx context.Context, project model.ProjectID, name string) (*entities.Workspace, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	namespace, err := r.getNamespaceFromProjectID(ctx, project)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.GetWorkspace(cc, namespace, name)
}

// CoreResyncWorkspace is the resolver for the core_resyncWorkspace field.
func (r *queryResolver) CoreResyncWorkspace(ctx context.Context, project model.ProjectID, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	namespace, err := r.getNamespaceFromProjectID(ctx, project)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.ResyncWorkspace(cc, namespace, name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreListEnvironments is the resolver for the core_listEnvironments field.
func (r *queryResolver) CoreListEnvironments(ctx context.Context, project model.ProjectID, search *model.SearchWorkspaces, pq *repos.CursorPagination) (*model.WorkspacePaginatedRecords, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
		if search.ProjectName != nil {
			filter["spec.projectName"] = *search.ProjectName
		}
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}
		if search.MarkedForDeletion != nil {
			filter["markedForDeletion"] = *search.MarkedForDeletion
		}
	}

	namespace, err := r.getNamespaceFromProjectID(ctx, project)
	if err != nil {
		return nil, errors.NewE(err)
	}

	pw, err := r.Domain.ListEnvironments(cc, namespace, filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	ee := make([]*model.WorkspaceEdge, len(pw.Edges))
	for i := range pw.Edges {
		ee[i] = &model.WorkspaceEdge{
			Node:   pw.Edges[i].Node,
			Cursor: pw.Edges[i].Cursor,
		}
	}

	m := model.WorkspacePaginatedRecords{
		Edges: ee,
		PageInfo: &model.PageInfo{
			EndCursor:       &pw.PageInfo.EndCursor,
			HasNextPage:     pw.PageInfo.HasNextPage,
			HasPreviousPage: pw.PageInfo.HasPrevPage,
			StartCursor:     &pw.PageInfo.StartCursor,
		},
		TotalCount: int(pw.TotalCount),
	}

	return &m, nil
}

// CoreGetEnvironment is the resolver for the core_getEnvironment field.
func (r *queryResolver) CoreGetEnvironment(ctx context.Context, project model.ProjectID, name string) (*entities.Workspace, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	namespace, err := r.getNamespaceFromProjectID(ctx, project)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.GetEnvironment(cc, namespace, name)
}

// CoreResyncEnvironment is the resolver for the core_resyncEnvironment field.
func (r *queryResolver) CoreResyncEnvironment(ctx context.Context, project model.ProjectID, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	namespace, err := r.getNamespaceFromProjectID(ctx, project)
	if err != nil {
		return false, errors.NewE(err)
	}

	if err := r.Domain.ResyncEnvironment(cc, namespace, name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreListApps is the resolver for the core_listApps field.
func (r *queryResolver) CoreListApps(ctx context.Context, project model.ProjectID, scope model.WorkspaceOrEnvID, search *model.SearchApps, pq *repos.CursorPagination) (*model.AppPaginatedRecords, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}
		if search.MarkedForDeletion != nil {
			filter["markedForDeletion"] = *search.MarkedForDeletion
		}
	}

	namespace, err := r.getNamespaceFromProjectAndScope(ctx, project, scope)
	if err != nil {
		return nil, errors.NewE(err)
	}

	pApps, err := r.Domain.ListApps(cc, namespace, filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	ae := make([]*model.AppEdge, len(pApps.Edges))
	for i := range pApps.Edges {
		ae[i] = &model.AppEdge{
			Node:   pApps.Edges[i].Node,
			Cursor: pApps.Edges[i].Cursor,
		}
	}

	m := model.AppPaginatedRecords{
		Edges: ae,
		PageInfo: &model.PageInfo{
			EndCursor:       &pApps.PageInfo.EndCursor,
			HasNextPage:     pApps.PageInfo.HasNextPage,
			HasPreviousPage: pApps.PageInfo.HasPrevPage,
			StartCursor:     &pApps.PageInfo.StartCursor,
		},
		TotalCount: int(pApps.TotalCount),
	}

	return &m, nil
}

// CoreGetApp is the resolver for the core_getApp field.
func (r *queryResolver) CoreGetApp(ctx context.Context, project model.ProjectID, scope model.WorkspaceOrEnvID, name string) (*entities.App, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	namespace, err := r.getNamespaceFromProjectAndScope(ctx, project, scope)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.GetApp(cc, namespace, name)
}

// CoreResyncApp is the resolver for the core_resyncApp field.
func (r *queryResolver) CoreResyncApp(ctx context.Context, project model.ProjectID, scope model.WorkspaceOrEnvID, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	namespace, err := r.getNamespaceFromProjectAndScope(ctx, project, scope)
	if err != nil {
		return false, errors.NewE(err)
	}

	if err := r.Domain.ResyncApp(cc, namespace, name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreListConfigs is the resolver for the core_listConfigs field.
func (r *queryResolver) CoreListConfigs(ctx context.Context, project model.ProjectID, scope model.WorkspaceOrEnvID, search *model.SearchConfigs, pq *repos.CursorPagination) (*model.ConfigPaginatedRecords, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}
		if search.MarkedForDeletion != nil {
			filter["markedForDeletion"] = *search.MarkedForDeletion
		}
	}

	namespace, err := r.getNamespaceFromProjectAndScope(ctx, project, scope)
	if err != nil {
		return nil, errors.NewE(err)
	}

	pConfigs, err := r.Domain.ListConfigs(cc, namespace, filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	ce := make([]*model.ConfigEdge, len(pConfigs.Edges))
	for i := range pConfigs.Edges {
		ce[i] = &model.ConfigEdge{
			Node:   pConfigs.Edges[i].Node,
			Cursor: pConfigs.Edges[i].Cursor,
		}
	}

	m := model.ConfigPaginatedRecords{
		Edges: ce,
		PageInfo: &model.PageInfo{
			EndCursor:       &pConfigs.PageInfo.EndCursor,
			HasNextPage:     pConfigs.PageInfo.HasNextPage,
			HasPreviousPage: pConfigs.PageInfo.HasPrevPage,
			StartCursor:     &pConfigs.PageInfo.StartCursor,
		},
		TotalCount: int(pConfigs.TotalCount),
	}

	return &m, nil
}

// CoreGetConfig is the resolver for the core_getConfig field.
func (r *queryResolver) CoreGetConfig(ctx context.Context, project model.ProjectID, scope model.WorkspaceOrEnvID, name string) (*entities.Config, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	namespace, err := r.getNamespaceFromProjectAndScope(ctx, project, scope)
	if err != nil {
		return nil, errors.NewE(err)
	}
	return r.Domain.GetConfig(cc, namespace, name)
}

// CoreResyncConfig is the resolver for the core_resyncConfig field.
func (r *queryResolver) CoreResyncConfig(ctx context.Context, project model.ProjectID, scope model.WorkspaceOrEnvID, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	namespace, err := r.getNamespaceFromProjectAndScope(ctx, project, scope)
	if err != nil {
		return false, errors.NewE(err)
	}
	if err := r.Domain.ResyncConfig(cc, namespace, name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreListSecrets is the resolver for the core_listSecrets field.
func (r *queryResolver) CoreListSecrets(ctx context.Context, project model.ProjectID, scope model.WorkspaceOrEnvID, search *model.SearchSecrets, pq *repos.CursorPagination) (*model.SecretPaginatedRecords, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}
		if search.MarkedForDeletion != nil {
			filter["markedForDeletion"] = *search.MarkedForDeletion
		}
	}

	namespace, err := r.getNamespaceFromProjectAndScope(ctx, project, scope)
	if err != nil {
		return nil, errors.NewE(err)
	}

	pSecrets, err := r.Domain.ListSecrets(cc, namespace, filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	ae := make([]*model.SecretEdge, len(pSecrets.Edges))
	for i := range pSecrets.Edges {
		ae[i] = &model.SecretEdge{
			Node:   pSecrets.Edges[i].Node,
			Cursor: pSecrets.Edges[i].Cursor,
		}
	}

	m := model.SecretPaginatedRecords{
		Edges: ae,
		PageInfo: &model.PageInfo{
			EndCursor:       &pSecrets.PageInfo.EndCursor,
			HasNextPage:     pSecrets.PageInfo.HasNextPage,
			HasPreviousPage: pSecrets.PageInfo.HasPrevPage,
			StartCursor:     &pSecrets.PageInfo.StartCursor,
		},
		TotalCount: int(pSecrets.TotalCount),
	}

	return &m, nil
}

// CoreGetSecret is the resolver for the core_getSecret field.
func (r *queryResolver) CoreGetSecret(ctx context.Context, project model.ProjectID, scope model.WorkspaceOrEnvID, name string) (*entities.Secret, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	namespace, err := r.getNamespaceFromProjectAndScope(ctx, project, scope)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.GetSecret(cc, namespace, name)
}

// CoreResyncSecret is the resolver for the core_resyncSecret field.
func (r *queryResolver) CoreResyncSecret(ctx context.Context, project model.ProjectID, scope model.WorkspaceOrEnvID, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	namespace, err := r.getNamespaceFromProjectAndScope(ctx, project, scope)
	if err != nil {
		return false, errors.NewE(err)
	}

	if err := r.Domain.ResyncSecret(cc, namespace, name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreListRouters is the resolver for the core_listRouters field.
func (r *queryResolver) CoreListRouters(ctx context.Context, project model.ProjectID, scope model.WorkspaceOrEnvID, search *model.SearchRouters, pq *repos.CursorPagination) (*model.RouterPaginatedRecords, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}
		if search.MarkedForDeletion != nil {
			filter["markedForDeletion"] = *search.MarkedForDeletion
		}
	}

	namespace, err := r.getNamespaceFromProjectAndScope(ctx, project, scope)
	if err != nil {
		return nil, errors.NewE(err)
	}

	pRouters, err := r.Domain.ListRouters(cc, namespace, filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	ae := make([]*model.RouterEdge, len(pRouters.Edges))
	for i := range pRouters.Edges {
		ae[i] = &model.RouterEdge{
			Node:   pRouters.Edges[i].Node,
			Cursor: pRouters.Edges[i].Cursor,
		}
	}

	m := model.RouterPaginatedRecords{
		Edges: ae,
		PageInfo: &model.PageInfo{
			EndCursor:       &pRouters.PageInfo.EndCursor,
			HasNextPage:     pRouters.PageInfo.HasNextPage,
			HasPreviousPage: pRouters.PageInfo.HasPrevPage,
			StartCursor:     &pRouters.PageInfo.StartCursor,
		},
		TotalCount: int(pRouters.TotalCount),
	}

	return &m, nil
}

// CoreGetRouter is the resolver for the core_getRouter field.
func (r *queryResolver) CoreGetRouter(ctx context.Context, project model.ProjectID, scope model.WorkspaceOrEnvID, name string) (*entities.Router, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	namespace, err := r.getNamespaceFromProjectAndScope(ctx, project, scope)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.GetRouter(cc, namespace, name)
}

// CoreResyncRouter is the resolver for the core_resyncRouter field.
func (r *queryResolver) CoreResyncRouter(ctx context.Context, project model.ProjectID, scope model.WorkspaceOrEnvID, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	namespace, err := r.getNamespaceFromProjectAndScope(ctx, project, scope)
	if err != nil {
		return false, errors.NewE(err)
	}

	if err := r.Domain.ResyncRouter(cc, namespace, name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreListManagedServiceTemplates is the resolver for the core_listManagedServiceTemplates field.
func (r *queryResolver) CoreListManagedServiceTemplates(ctx context.Context) ([]*entities.MsvcTemplate, error) {
	return r.Domain.ListManagedSvcTemplates()
}

// CoreGetManagedServiceTemplate is the resolver for the core_getManagedServiceTemplate field.
func (r *queryResolver) CoreGetManagedServiceTemplate(ctx context.Context, category string, name string) (*entities.MsvcTemplateEntry, error) {
	return r.Domain.GetManagedSvcTemplate(category, name)
}

// CoreListManagedServices is the resolver for the core_listManagedServices field.
func (r *queryResolver) CoreListManagedServices(ctx context.Context, project model.ProjectID, scope model.WorkspaceOrEnvID, search *model.SearchManagedServices, pq *repos.CursorPagination) (*model.ManagedServicePaginatedRecords, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}
		if search.MarkedForDeletion != nil {
			filter["markedForDeletion"] = *search.MarkedForDeletion
		}
	}

	namespace, err := r.getNamespaceFromProjectAndScope(ctx, project, scope)
	if err != nil {
		return nil, errors.NewE(err)
	}

	pMsvcs, err := r.Domain.ListManagedServices(cc, namespace, filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	msvcEdges := make([]*model.ManagedServiceEdge, len(pMsvcs.Edges))
	for i := range pMsvcs.Edges {
		msvcEdges[i] = &model.ManagedServiceEdge{
			Node:   pMsvcs.Edges[i].Node,
			Cursor: pMsvcs.Edges[i].Cursor,
		}
	}

	m := model.ManagedServicePaginatedRecords{
		Edges: msvcEdges,
		PageInfo: &model.PageInfo{
			EndCursor:       &pMsvcs.PageInfo.EndCursor,
			HasNextPage:     pMsvcs.PageInfo.HasNextPage,
			HasPreviousPage: pMsvcs.PageInfo.HasPrevPage,
			StartCursor:     &pMsvcs.PageInfo.StartCursor,
		},
		TotalCount: int(pMsvcs.TotalCount),
	}

	return &m, nil
}

// CoreGetManagedService is the resolver for the core_getManagedService field.
func (r *queryResolver) CoreGetManagedService(ctx context.Context, project model.ProjectID, scope model.WorkspaceOrEnvID, name string) (*entities.ManagedService, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	namespace, err := r.getNamespaceFromProjectAndScope(ctx, project, scope)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.GetManagedService(cc, namespace, name)
}

// CoreResyncManagedService is the resolver for the core_resyncManagedService field.
func (r *queryResolver) CoreResyncManagedService(ctx context.Context, project model.ProjectID, scope model.WorkspaceOrEnvID, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	namespace, err := r.getNamespaceFromProjectAndScope(ctx, project, scope)
	if err != nil {
		return false, errors.NewE(err)
	}

	if err := r.Domain.ResyncManagedService(cc, namespace, name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// CoreListManagedResources is the resolver for the core_listManagedResources field.
func (r *queryResolver) CoreListManagedResources(ctx context.Context, project model.ProjectID, scope model.WorkspaceOrEnvID, search *model.SearchManagedResources, pq *repos.CursorPagination) (*model.ManagedResourcePaginatedRecords, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}
		if search.MarkedForDeletion != nil {
			filter["markedForDeletion"] = *search.MarkedForDeletion
		}

		if search.ManagedServiceName != nil {
			filter["spec.msvcRef.name"] = *search.ManagedServiceName
		}
	}

	namespace, err := r.getNamespaceFromProjectAndScope(ctx, project, scope)
	if err != nil {
		return nil, errors.NewE(err)
	}

	pApps, err := r.Domain.ListManagedResources(cc, namespace, filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, errors.NewE(err)
	}

	ae := make([]*model.ManagedResourceEdge, len(pApps.Edges))
	for i := range pApps.Edges {
		ae[i] = &model.ManagedResourceEdge{
			Node:   pApps.Edges[i].Node,
			Cursor: pApps.Edges[i].Cursor,
		}
	}

	m := model.ManagedResourcePaginatedRecords{
		Edges: ae,
		PageInfo: &model.PageInfo{
			EndCursor:       &pApps.PageInfo.EndCursor,
			HasNextPage:     pApps.PageInfo.HasNextPage,
			HasPreviousPage: pApps.PageInfo.HasPrevPage,
			StartCursor:     &pApps.PageInfo.StartCursor,
		},
		TotalCount: int(pApps.TotalCount),
	}

	return &m, nil
}

// CoreGetManagedResource is the resolver for the core_getManagedResource field.
func (r *queryResolver) CoreGetManagedResource(ctx context.Context, project model.ProjectID, scope model.WorkspaceOrEnvID, name string) (*entities.ManagedResource, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return nil, errors.NewE(err)
	}
	namespace, err := r.getNamespaceFromProjectAndScope(ctx, project, scope)
	if err != nil {
		return nil, errors.NewE(err)
	}

	return r.Domain.GetManagedResource(cc, namespace, name)
}

// CoreResyncManagedResource is the resolver for the core_resyncManagedResource field.
func (r *queryResolver) CoreResyncManagedResource(ctx context.Context, project model.ProjectID, scope model.WorkspaceOrEnvID, name string) (bool, error) {
	cc, err := toConsoleContext(ctx)
	if err != nil {
		return false, errors.NewE(err)
	}
	namespace, err := r.getNamespaceFromProjectAndScope(ctx, project, scope)
	if err != nil {
		return false, errors.NewE(err)
	}

	if err := r.Domain.ResyncManagedResource(cc, namespace, name); err != nil {
		return false, errors.NewE(err)
	}
	return true, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
