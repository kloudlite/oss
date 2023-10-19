package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"fmt"

	"kloudlite.io/apps/infra/internal/app/graph/generated"
	"kloudlite.io/apps/infra/internal/app/graph/model"
	"kloudlite.io/apps/infra/internal/domain"
	"kloudlite.io/apps/infra/internal/entities"
	fn "kloudlite.io/pkg/functions"
	"kloudlite.io/pkg/repos"
)

// InfraCreateCluster is the resolver for the infra_createCluster field.
func (r *mutationResolver) InfraCreateCluster(ctx context.Context, cluster entities.Cluster) (*entities.Cluster, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.Domain.CreateCluster(ictx, cluster)
}

// InfraUpdateCluster is the resolver for the infra_updateCluster field.
func (r *mutationResolver) InfraUpdateCluster(ctx context.Context, cluster entities.Cluster) (*entities.Cluster, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.UpdateCluster(ictx, cluster)
}

// InfraDeleteCluster is the resolver for the infra_deleteCluster field.
func (r *mutationResolver) InfraDeleteCluster(ctx context.Context, name string) (bool, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return false, err
	}
	if err := r.Domain.DeleteCluster(ictx, name); err != nil {
		return false, err
	}
	return true, nil
}

// InfraCreateBYOCCluster is the resolver for the infra_createBYOCCluster field.
func (r *mutationResolver) InfraCreateBYOCCluster(ctx context.Context, byocCluster entities.BYOCCluster) (*entities.BYOCCluster, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.CreateBYOCCluster(ictx, byocCluster)
}

// InfraUpdateBYOCCluster is the resolver for the infra_updateBYOCCluster field.
func (r *mutationResolver) InfraUpdateBYOCCluster(ctx context.Context, byocCluster entities.BYOCCluster) (*entities.BYOCCluster, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.UpdateBYOCCluster(ictx, byocCluster)
}

// InfraDeleteBYOCCluster is the resolver for the infra_deleteBYOCCluster field.
func (r *mutationResolver) InfraDeleteBYOCCluster(ctx context.Context, name string) (bool, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return false, err
	}

	if err := r.Domain.DeleteBYOCCluster(ictx, name); err != nil {
		return false, err
	}
	return true, nil
}

// InfraCreateProviderSecret is the resolver for the infra_createProviderSecret field.
func (r *mutationResolver) InfraCreateProviderSecret(ctx context.Context, secret entities.CloudProviderSecret) (*entities.CloudProviderSecret, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.CreateProviderSecret(ictx, secret)
}

// InfraUpdateProviderSecret is the resolver for the infra_updateProviderSecret field.
func (r *mutationResolver) InfraUpdateProviderSecret(ctx context.Context, secret entities.CloudProviderSecret) (*entities.CloudProviderSecret, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.UpdateProviderSecret(ictx, secret)
}

// InfraDeleteProviderSecret is the resolver for the infra_deleteProviderSecret field.
func (r *mutationResolver) InfraDeleteProviderSecret(ctx context.Context, secretName string) (bool, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return false, err
	}

	if err := r.Domain.DeleteProviderSecret(ictx, secretName); err != nil {
		return false, err
	}
	return true, nil
}

// InfraCreateDomainEntry is the resolver for the infra_createDomainEntry field.
func (r *mutationResolver) InfraCreateDomainEntry(ctx context.Context, domainEntry entities.DomainEntry) (*entities.DomainEntry, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.Domain.CreateDomainEntry(ictx, domainEntry)
}

// InfraUpdateDomainEntry is the resolver for the infra_updateDomainEntry field.
func (r *mutationResolver) InfraUpdateDomainEntry(ctx context.Context, domainEntry entities.DomainEntry) (*entities.DomainEntry, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.Domain.UpdateDomainEntry(ictx, domainEntry)
}

// InfraDeleteDomainEntry is the resolver for the infra_deleteDomainEntry field.
func (r *mutationResolver) InfraDeleteDomainEntry(ctx context.Context, domainName string) (bool, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return false, err
	}
	if err := r.Domain.DeleteDomainEntry(ictx, domainName); err != nil {
		return false, err
	}
	return true, nil
}

// InfraCreateNodePool is the resolver for the infra_createNodePool field.
func (r *mutationResolver) InfraCreateNodePool(ctx context.Context, clusterName string, pool entities.NodePool) (*entities.NodePool, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.CreateNodePool(ictx, clusterName, pool)
}

// InfraUpdateNodePool is the resolver for the infra_updateNodePool field.
func (r *mutationResolver) InfraUpdateNodePool(ctx context.Context, clusterName string, pool entities.NodePool) (*entities.NodePool, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.UpdateNodePool(ictx, clusterName, pool)
}

// InfraDeleteNodePool is the resolver for the infra_deleteNodePool field.
func (r *mutationResolver) InfraDeleteNodePool(ctx context.Context, clusterName string, poolName string) (bool, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return false, err
	}

	if err := r.Domain.DeleteNodePool(ictx, clusterName, poolName); err != nil {
		return false, err
	}
	return true, nil
}

// CoreCreateVPNDevice is the resolver for the core_createVPNDevice field.
func (r *mutationResolver) CoreCreateVPNDevice(ctx context.Context, clusterName string, vpnDevice entities.VPNDevice) (*entities.VPNDevice, error) {
	cc, err := toInfraContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.CreateVPNDevice(cc, clusterName, vpnDevice)
}

// CoreUpdateVPNDevice is the resolver for the core_updateVPNDevice field.
func (r *mutationResolver) CoreUpdateVPNDevice(ctx context.Context, clusterName string, vpnDevice entities.VPNDevice) (*entities.VPNDevice, error) {
	cc, err := toInfraContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.Domain.UpdateVPNDevice(cc, clusterName, vpnDevice)
}

// CoreDeleteVPNDevice is the resolver for the core_deleteVPNDevice field.
func (r *mutationResolver) CoreDeleteVPNDevice(ctx context.Context, clusterName string, deviceName string) (bool, error) {
	cc, err := toInfraContext(ctx)
	if err != nil {
		return false, err
	}
	if err := r.Domain.DeleteVPNDevice(cc, clusterName, deviceName); err != nil {
		return false, err
	}
	return true, nil
}

// InfraCheckNameAvailability is the resolver for the infra_checkNameAvailability field.
func (r *queryResolver) InfraCheckNameAvailability(ctx context.Context, resType domain.ResType, clusterName *string, name string) (*domain.CheckNameAvailabilityOutput, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.CheckNameAvailability(ictx, resType, clusterName, name)
}

// InfraListClusters is the resolver for the infra_listClusters field.
func (r *queryResolver) InfraListClusters(ctx context.Context, search *model.SearchCluster, pagination *repos.CursorPagination) (*model.ClusterPaginatedRecords, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, err
	}

	if pagination == nil {
		pagination = &repos.DefaultCursorPagination
	}

	filter := map[string]repos.MatchFilter{}

	if search != nil {
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}

		if search.CloudProviderName != nil {
			filter["spec.cloudProvider"] = *search.CloudProviderName
		}

		if search.Region != nil {
			filter["spec.region"] = *search.Region
		}

		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
	}

	pClusters, err := r.Domain.ListClusters(ictx, filter, *pagination)
	if err != nil {
		return nil, err
	}

	ce := make([]*model.ClusterEdge, len(pClusters.Edges))
	for i := range pClusters.Edges {
		ce[i] = &model.ClusterEdge{
			Node:   pClusters.Edges[i].Node,
			Cursor: pClusters.Edges[i].Cursor,
		}
	}

	m := model.ClusterPaginatedRecords{
		Edges: ce,
		PageInfo: &model.PageInfo{
			EndCursor:       &pClusters.PageInfo.EndCursor,
			HasNextPage:     pClusters.PageInfo.HasNextPage,
			HasPreviousPage: pClusters.PageInfo.HasPrevPage,
			StartCursor:     &pClusters.PageInfo.StartCursor,
		},
		TotalCount: int(pClusters.TotalCount),
	}

	return &m, nil
}

// InfraGetCluster is the resolver for the infra_getCluster field.
func (r *queryResolver) InfraGetCluster(ctx context.Context, name string) (*entities.Cluster, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.GetCluster(ictx, name)
}

// InfraListBYOCClusters is the resolver for the infra_listBYOCClusters field.
func (r *queryResolver) InfraListBYOCClusters(ctx context.Context, search *model.SearchCluster, pagination *repos.CursorPagination) (*model.BYOCClusterPaginatedRecords, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, err
	}

	if pagination == nil {
		pagination = &repos.DefaultCursorPagination
	}

	filter := map[string]repos.MatchFilter{}
	if search != nil {
		if search.IsReady != nil {
			filter["status.isReady"] = *search.IsReady
		}

		if search.CloudProviderName != nil {
			filter["spec.cloudProvider"] = *search.CloudProviderName
		}

		if search.Region != nil {
			filter["spec.region"] = *search.Region
		}

		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
	}

	pClusters, err := r.Domain.ListBYOCClusters(ictx, filter, *pagination)
	if err != nil {
		return nil, err
	}

	bce := make([]*model.BYOCClusterEdge, len(pClusters.Edges))
	for i := range pClusters.Edges {
		bce[i] = &model.BYOCClusterEdge{
			Node:   pClusters.Edges[i].Node,
			Cursor: pClusters.Edges[i].Cursor,
		}
	}

	m := model.BYOCClusterPaginatedRecords{
		Edges: bce,
		PageInfo: &model.PageInfo{
			EndCursor:       &pClusters.PageInfo.EndCursor,
			HasNextPage:     pClusters.PageInfo.HasNextPage,
			HasPreviousPage: pClusters.PageInfo.HasPrevPage,
			StartCursor:     &pClusters.PageInfo.StartCursor,
		},
		TotalCount: int(pClusters.TotalCount),
	}

	return &m, nil
}

// InfraGetBYOCCluster is the resolver for the infra_getBYOCCluster field.
func (r *queryResolver) InfraGetBYOCCluster(ctx context.Context, name string) (*entities.BYOCCluster, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.GetBYOCCluster(ictx, name)
}

// InfraListNodePools is the resolver for the infra_listNodePools field.
func (r *queryResolver) InfraListNodePools(ctx context.Context, clusterName string, search *model.SearchNodepool, pagination *repos.CursorPagination) (*model.NodePoolPaginatedRecords, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, err
	}

	if pagination == nil {
		pagination = &repos.DefaultCursorPagination
	}

	filter := map[string]repos.MatchFilter{}

	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}
	}

	pNodePools, err := r.Domain.ListNodePools(ictx, clusterName, filter, *pagination)
	if err != nil {
		return nil, err
	}

	pe := make([]*model.NodePoolEdge, len(pNodePools.Edges))
	for i := range pNodePools.Edges {
		pe[i] = &model.NodePoolEdge{
			Node:   pNodePools.Edges[i].Node,
			Cursor: pNodePools.Edges[i].Cursor,
		}
	}

	m := model.NodePoolPaginatedRecords{
		Edges: pe,
		PageInfo: &model.PageInfo{
			EndCursor:       &pNodePools.PageInfo.EndCursor,
			HasNextPage:     pNodePools.PageInfo.HasNextPage,
			HasPreviousPage: pNodePools.PageInfo.HasPrevPage,
			StartCursor:     &pNodePools.PageInfo.StartCursor,
		},
		TotalCount: int(pNodePools.TotalCount),
	}

	return &m, nil
}

// InfraGetNodePool is the resolver for the infra_getNodePool field.
func (r *queryResolver) InfraGetNodePool(ctx context.Context, clusterName string, poolName string) (*entities.NodePool, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.GetNodePool(ictx, clusterName, poolName)
}

// InfraListProviderSecrets is the resolver for the infra_listProviderSecrets field.
func (r *queryResolver) InfraListProviderSecrets(ctx context.Context, search *model.SearchProviderSecret, pagination *repos.CursorPagination) (*model.CloudProviderSecretPaginatedRecords, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, err
	}

	if pagination == nil {
		pagination = &repos.DefaultCursorPagination
	}

	filter := map[string]repos.MatchFilter{}

	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}

		if search.CloudProviderName != nil {
			filter["cloudProviderName"] = *search.CloudProviderName
		}
	}

	pSecrets, err := r.Domain.ListProviderSecrets(ictx, filter, *pagination)
	if err != nil {
		return nil, err
	}

	pe := make([]*model.CloudProviderSecretEdge, len(pSecrets.Edges))
	for i := range pSecrets.Edges {
		pe[i] = &model.CloudProviderSecretEdge{
			Node:   pSecrets.Edges[i].Node,
			Cursor: pSecrets.Edges[i].Cursor,
		}
	}

	m := model.CloudProviderSecretPaginatedRecords{
		Edges: pe,
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

// InfraGetProviderSecret is the resolver for the infra_getProviderSecret field.
func (r *queryResolver) InfraGetProviderSecret(ctx context.Context, name string) (*entities.CloudProviderSecret, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.GetProviderSecret(ictx, name)
}

// InfraListDomainEntries is the resolver for the infra_listDomainEntries field.
func (r *queryResolver) InfraListDomainEntries(ctx context.Context, search *model.SearchDomainEntry, pagination *repos.CursorPagination) (*model.DomainEntryPaginatedRecords, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, err
	}

	filter := map[string]repos.MatchFilter{}

	if search != nil {
		if search.Text != nil {
			filter["metadata.name"] = *search.Text
		}

		if search.ClusterName != nil {
			filter["spec.clusterName"] = *search.ClusterName
		}
	}

	dEntries, err := r.Domain.ListDomainEntries(ictx, filter, fn.DefaultIfNil(pagination, repos.DefaultCursorPagination))
	if err != nil {
		return nil, err
	}

	edges := make([]*model.DomainEntryEdge, len(dEntries.Edges))
	for i := range dEntries.Edges {
		edges[i] = &model.DomainEntryEdge{
			Node:   dEntries.Edges[i].Node,
			Cursor: dEntries.Edges[i].Cursor,
		}
	}

	m := model.DomainEntryPaginatedRecords{
		Edges: edges,
		PageInfo: &model.PageInfo{
			EndCursor:       &dEntries.PageInfo.EndCursor,
			HasNextPage:     dEntries.PageInfo.HasNextPage,
			HasPreviousPage: dEntries.PageInfo.HasPrevPage,
			StartCursor:     &dEntries.PageInfo.StartCursor,
		},
		TotalCount: int(dEntries.TotalCount),
	}

	return &m, nil
}

// InfraGetDomainEntry is the resolver for the infra_getDomainEntry field.
func (r *queryResolver) InfraGetDomainEntry(ctx context.Context, domainName string) (*entities.DomainEntry, error) {
	ictx, err := toInfraContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.Domain.GetDomainEntry(ictx, domainName)
}

// CoreListVPNDevices is the resolver for the core_listVPNDevices field.
func (r *queryResolver) CoreListVPNDevices(ctx context.Context, clusterName *string, search *model.SearchVPNDevices, pq *repos.CursorPagination) (*model.VPNDevicePaginatedRecords, error) {
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

	cc, err := toInfraContext(ctx)
	if err != nil {
		if cc.AccountName == "" {
			return nil, err
		}
	}

	devices, err := r.Domain.ListVPNDevices(cc, cc.AccountName, clusterName, filter, fn.DefaultIfNil(pq, repos.DefaultCursorPagination))
	if err != nil {
		return nil, err
	}

	ve := make([]*model.VPNDeviceEdge, len(devices.Edges))
	for i := range devices.Edges {
		ve[i] = &model.VPNDeviceEdge{
			Node:   devices.Edges[i].Node,
			Cursor: devices.Edges[i].Cursor,
		}
	}

	m := model.VPNDevicePaginatedRecords{
		Edges: ve,
		PageInfo: &model.PageInfo{
			EndCursor:       &devices.PageInfo.EndCursor,
			HasNextPage:     devices.PageInfo.HasNextPage,
			HasPreviousPage: devices.PageInfo.HasPrevPage,
			StartCursor:     &devices.PageInfo.StartCursor,
		},
		TotalCount: int(devices.TotalCount),
	}

	return &m, nil
}

// CoreGetVPNDevice is the resolver for the core_getVPNDevice field.
func (r *queryResolver) CoreGetVPNDevice(ctx context.Context, clusterName string, name string) (*entities.VPNDevice, error) {
	cc, err := toInfraContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.Domain.GetVPNDevice(cc, "", name)
}

// CoreGetVPNDeviceConfig is the resolver for the core_getVPNDeviceConfig field.
func (r *queryResolver) CoreGetVPNDeviceConfig(ctx context.Context, clusterName string, name string) (string, error) {
	panic(fmt.Errorf("not implemented: CoreGetVPNDeviceConfig - core_getVPNDeviceConfig"))
}

// WgConfig is the resolver for the wgConfig field.
func (r *vPNDeviceResolver) WgConfig(ctx context.Context, obj *entities.VPNDevice) (*string, error) {
	panic(fmt.Errorf("not implemented: WgConfig - wgConfig"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
