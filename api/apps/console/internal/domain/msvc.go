package domain

import (
	"fmt"
	"time"

	"kloudlite.io/apps/console/internal/domain/entities"
	"kloudlite.io/pkg/repos"
	t "kloudlite.io/pkg/types"
)

func (d *domain) ListManagedServices(ctx ConsoleContext, namespace string) ([]*entities.MSvc, error) {
	if err := d.canReadResourcesInProject(ctx, namespace); err != nil {
		return nil, err
	}
	return d.msvcRepo.Find(ctx, repos.Query{Filter: repos.Filter{
		"accountName":        ctx.AccountName,
		"clusterName":        ctx.ClusterName,
		"metadata.namespace": namespace,
	}})
}

func (d *domain) findMSvc(ctx ConsoleContext, namespace string, name string) (*entities.MSvc, error) {
	mres, err := d.msvcRepo.FindOne(ctx, repos.Filter{
		"accountName":        ctx.AccountName,
		"clusterName":        ctx.ClusterName,
		"metadata.namespace": namespace,
		"metadata.name":      name,
	})
	if err != nil {
		return nil, err
	}
	if mres == nil {
		return nil, fmt.Errorf("no secret with name=%q,namespace=%q found", name, namespace)
	}
	return mres, nil
}

func (d *domain) GetManagedService(ctx ConsoleContext, namespace string, name string) (*entities.MSvc, error) {
	if err := d.canReadResourcesInProject(ctx, namespace); err != nil {
		return nil, err
	}
	return d.findMSvc(ctx, namespace, name)
}

// mutations

func (d *domain) CreateManagedService(ctx ConsoleContext, msvc entities.MSvc) (*entities.MSvc, error) {
	if err := d.canMutateResourcesInProject(ctx, msvc.Namespace); err != nil {
		return nil, err
	}

	msvc.EnsureGVK()
	if err := d.k8sExtendedClient.ValidateStruct(ctx, &msvc.ManagedService); err != nil {
		return nil, err
	}

	msvc.AccountName = ctx.AccountName
	msvc.ClusterName = ctx.ClusterName
	msvc.Generation = 1
	msvc.SyncStatus = t.GenSyncStatus(t.SyncActionApply, msvc.Generation)

	m, err := d.msvcRepo.Create(ctx, &msvc)
	if err != nil {
		if d.msvcRepo.ErrAlreadyExists(err) {
			return nil, fmt.Errorf("msvc with name=%q, namespace=%q already exists", msvc.Name, msvc.Namespace)
		}
		return nil, err
	}

	if err := d.applyK8sResource(ctx, &m.ManagedService); err != nil {
		return m, err
	}

	return m, nil
}

func (d *domain) UpdateManagedService(ctx ConsoleContext, msvc entities.MSvc) (*entities.MSvc, error) {
	if err := d.canMutateResourcesInProject(ctx, msvc.Namespace); err != nil {
		return nil, err
	}

	msvc.EnsureGVK()
	if err := d.k8sExtendedClient.ValidateStruct(ctx, &msvc.ManagedService); err != nil {
		return nil, err
	}

	m, err := d.findMSvc(ctx, msvc.Namespace, msvc.Name)
	if err != nil {
		return nil, err
	}

	m.Spec = msvc.Spec
	m.Generation += 1
	m.SyncStatus = t.GenSyncStatus(t.SyncActionApply, m.Generation)

	upMSvc, err := d.msvcRepo.UpdateById(ctx, m.Id, m)
	if err != nil {
		return nil, err
	}

	if err := d.applyK8sResource(ctx, &upMSvc.ManagedService); err != nil {
		return upMSvc, err
	}

	return upMSvc, nil
}

func (d *domain) DeleteManagedService(ctx ConsoleContext, namespace string, name string) error {
	if err := d.canMutateResourcesInProject(ctx, namespace); err != nil {
		return err
	}
	m, err := d.findMSvc(ctx, namespace, name)
	if err != nil {
		return err
	}

	m.SyncStatus = t.GenSyncStatus(t.SyncActionDelete, m.Generation)
	if _, err := d.msvcRepo.UpdateById(ctx, m.Id, m); err != nil {
		return err
	}

	return d.deleteK8sResource(ctx, &m.ManagedService)
}

func (d *domain) OnDeleteManagedServiceMessage(ctx ConsoleContext, msvc entities.MSvc) error {
	m, err := d.findMSvc(ctx, msvc.Namespace, msvc.Name)
	if err != nil {
		return err
	}

	return d.msvcRepo.DeleteById(ctx, m.Id)
}

func (d *domain) OnUpdateManagedServiceMessage(ctx ConsoleContext, msvc entities.MSvc) error {
	m, err := d.findMSvc(ctx, msvc.Namespace, msvc.Name)
	if err != nil {
		return err
	}

	m.Status = msvc.Status
	m.SyncStatus.LastSyncedAt = time.Now()
	m.SyncStatus.Generation = msvc.Generation
	m.SyncStatus.State = t.ParseSyncState(msvc.Status.IsReady)

	_, err = d.msvcRepo.UpdateById(ctx, m.Id, m)
	return err
}

func (d *domain) OnApplyManagedServiceError(ctx ConsoleContext, err error, namespace, name string) error {
	m, err2 := d.findMSvc(ctx, namespace, name)
	if err2 != nil {
		return err2
	}

	m.SyncStatus.Error = err.Error()
	_, err2 = d.msvcRepo.UpdateById(ctx, m.Id, m)
	return err2
}

func (d *domain) ResyncManagedService(ctx ConsoleContext, namespace, name string) error {
	c, err := d.findMSvc(ctx, namespace, name)
	if err != nil {
		return err
	}
	return d.applyK8sResource(ctx, &c.ManagedService)
}
