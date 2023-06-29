package domain

import (
	"fmt"
	"time"

	"kloudlite.io/apps/infra/internal/domain/entities"
	"kloudlite.io/common"
	"kloudlite.io/pkg/repos"
	t "kloudlite.io/pkg/types"
)

func (d *domain) CreateCluster(ctx InfraContext, cluster entities.Cluster) (*entities.Cluster, error) {
	cluster.EnsureGVK()
	if err := d.k8sExtendedClient.ValidateStruct(ctx, &cluster.Cluster); err != nil {
		return nil, err
	}

	cluster.AccountName = ctx.AccountName
	cluster.SyncStatus = t.GetSyncStatusForCreation()

	nCluster, err := d.clusterRepo.Create(ctx, &cluster)
	if err != nil {
		if d.clusterRepo.ErrAlreadyExists(err) {
			return nil, fmt.Errorf("cluster with name %q already exists", cluster.Name)
		}
		return nil, err
	}

	if err := d.applyK8sResource(ctx, &nCluster.Cluster); err != nil {
		return nil, err
	}

	return nCluster, nil
}

func (d *domain) ListClusters(ctx InfraContext, pagination t.CursorPagination) (*repos.PaginatedRecord[*entities.Cluster], error) {
	return d.clusterRepo.FindPaginated(ctx, repos.Filter{
		"accountName": ctx.AccountName,
	}, pagination)
}

func (d *domain) GetCluster(ctx InfraContext, name string) (*entities.Cluster, error) {
	return d.clusterRepo.FindOne(ctx, repos.Filter{
		"accountName":   ctx.AccountName,
		"metadata.name": name,
	})
}

func (d *domain) UpdateCluster(ctx InfraContext, cluster entities.Cluster) (*entities.Cluster, error) {
	cluster.EnsureGVK()
	clus, err := d.findCluster(ctx, cluster.Name)
	if err != nil {
		return nil, err
	}

	clus.Cluster = cluster.Cluster
	clus.SyncStatus = t.GetSyncStatusForUpdation(clus.Generation + 1)

	uCluster, err := d.clusterRepo.UpdateById(ctx, clus.Id, clus)
	if err != nil {
		return nil, err
	}

	if err := d.applyK8sResource(ctx, &uCluster.Cluster); err != nil {
		return nil, err
	}

	return uCluster, nil
}

func (d *domain) DeleteCluster(ctx InfraContext, name string) error {
	c, err := d.findCluster(ctx, name)
	if err != nil {
		return err
	}

	c.SyncStatus = t.GetSyncStatusForDeletion(c.Generation)
	upC, err := d.clusterRepo.UpdateById(ctx, c.Id, c)
	if err != nil {
		return err
	}
	return d.deleteK8sResource(ctx, &upC.Cluster)
}

func (d *domain) OnDeleteClusterMessage(ctx InfraContext, cluster entities.Cluster) error {
	return d.clusterRepo.DeleteOne(ctx, repos.Filter{
		"accountName":   ctx.AccountName,
		"metadata.name": cluster.Name,
	})
}

func (d *domain) OnUpdateClusterMessage(ctx InfraContext, cluster entities.Cluster) error {
	c, err := d.findCluster(ctx, cluster.Name)
	if err != nil {
		return err
	}

	c.Cluster = cluster.Cluster
	c.SyncStatus.LastSyncedAt = time.Now()
	c.SyncStatus.State = t.ParseSyncState(c.Status.IsReady)

	_, err = d.clusterRepo.UpdateById(ctx, c.Id, c)
	return err
}

func (d *domain) findCluster(ctx InfraContext, clusterName string) (*entities.Cluster, error) {
	cluster, err := d.clusterRepo.FindOne(ctx, repos.Filter{
		"accountName":   ctx.AccountName,
		"metadata.name": clusterName,
	})
	if err != nil {
		return nil, err
	}
	if cluster == nil {
		return nil, fmt.Errorf("cluster with name %q not found", clusterName)
	}
	return cluster, nil
}

func (d *domain) findBYOCCluster(ctx InfraContext, clusterName string) (*entities.BYOCCluster, error) {
	cluster, err := d.byocClusterRepo.FindOne(ctx, repos.Filter{
		"spec.accountName": ctx.AccountName,
		"metadata.name":    clusterName,
	})
	if err != nil {
		return nil, err
	}
	if cluster == nil {
		return nil, fmt.Errorf("cluster with name %q not found", clusterName)
	}
	return cluster, nil
}

func (d *domain) CreateBYOCCluster(ctx InfraContext, cluster entities.BYOCCluster) (*entities.BYOCCluster, error) {
	cluster.EnsureGVK()
	cluster.Spec.IncomingKafkaTopic = common.GetKafkaTopicName(ctx.AccountName, cluster.Name)

	if err := d.k8sExtendedClient.ValidateStruct(ctx, &cluster.BYOC); err != nil {
		return nil, err
	}

	cluster.IsConnected = false
	cluster.Spec.AccountName = ctx.AccountName
	cluster.SyncStatus = t.GetSyncStatusForCreation()

	nCluster, err := d.byocClusterRepo.Create(ctx, &cluster)
	if err != nil {
		if d.clusterRepo.ErrAlreadyExists(err) {
			return nil, fmt.Errorf("cluster with name %q already exists", cluster.Name)
		}
	}

	if err := d.applyK8sResource(ctx, &nCluster.BYOC); err != nil {
		return nil, err
	}

	return nCluster, nil
}

func (d *domain) ListBYOCClusters(ctx InfraContext, pagination t.CursorPagination) (*repos.PaginatedRecord[*entities.BYOCCluster], error) {
	return d.byocClusterRepo.FindPaginated(ctx, repos.Filter{
		"spec.accountName": ctx.AccountName,
	}, pagination)
}

func (d *domain) GetBYOCCluster(ctx InfraContext, name string) (*entities.BYOCCluster, error) {
	return d.findBYOCCluster(ctx, name)
}

func (d *domain) UpdateBYOCCluster(ctx InfraContext, cluster entities.BYOCCluster) (*entities.BYOCCluster, error) {
	cluster.EnsureGVK()
	c, err := d.findBYOCCluster(ctx, cluster.Name)
	if err != nil {
		return nil, err
	}
	c.BYOC = cluster.BYOC
	c.SyncStatus = t.GetSyncStatusForUpdation(c.Generation + 1)

	// c.Spec.AccountName = ctx.AccountName
	// c.Spec.Region = cluster.Spec.Region
	// c.Spec.Provider = cluster.Spec.Provider
	uCluster, err := d.byocClusterRepo.UpdateOne(ctx, repos.Filter{"metadata.name": cluster.Name}, c)
	if err != nil {
		return nil, err
	}

	if err := d.applyK8sResource(ctx, &uCluster.BYOC); err != nil {
		return nil, err
	}

	return uCluster, nil
}

func (d *domain) DeleteBYOCCluster(ctx InfraContext, name string) error {
	clus, err := d.findBYOCCluster(ctx, name)
	if err != nil {
		return err
	}

	clus.SyncStatus = t.GetSyncStatusForDeletion(clus.Generation)
	upC, err := d.byocClusterRepo.UpdateById(ctx, clus.Id, clus)
	if err != nil {
		return err
	}
	return d.deleteK8sResource(ctx, &upC.BYOC)
}

func (d *domain) OnDeleteBYOCClusterMessage(ctx InfraContext, cluster entities.BYOCCluster) error {
	return d.clusterRepo.DeleteOne(ctx, repos.Filter{
		"spec.accountName": ctx.AccountName,
		"metadata.name":    cluster.Name,
	})
}

func (d *domain) OnBYOCClusterHelmUpdates(ctx InfraContext, cluster entities.BYOCCluster) error {
	c, err := d.findBYOCCluster(ctx, cluster.Name)
	if err != nil {
		return err
	}

	_, err = d.byocClusterRepo.UpdateById(ctx, c.Id, &cluster)
	if err != nil {
		return err
	}
	return nil
}

// func (d *domain) OnUpdateBYOCClusterMessage(ctx InfraContext, cluster entities.BYOCCluster) error {
// 	d.findBYOCCluster(ctx, cluster.Name)
// }
