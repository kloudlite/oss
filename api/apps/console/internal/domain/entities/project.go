package entities

import (
	crdsv1 "github.com/kloudlite/operator/apis/crds/v1"
	"kloudlite.io/pkg/repos"
	t "kloudlite.io/pkg/types"
)

type Project struct {
	repos.BaseEntity `json:",inline"`
	crdsv1.Project   `json:",inline"`
	AccountName      string       `json:"accountName"`
	ClusterName      string       `json:"clusterName"`
	SyncStatus       t.SyncStatus `json:"syncStatus"`
}

var ProjectIndexes = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "metadata.name", Value: repos.IndexAsc},
			{Key: "clusterName", Value: repos.IndexAsc},
			{Key: "accountName", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}

type Environment struct {
	repos.BaseEntity `json:",inline"`
	crdsv1.Env       `json:",inline"`
	AccountName      string       `json:"accountName"`
	ClusterName      string       `json:"clusterName"`
	SyncStatus       t.SyncStatus `json:"syncStatus"`
}

var EnvironmentIndexes = []repos.IndexField{
	{
		Field: []repos.IndexKey{
			{Key: "id", Value: repos.IndexAsc},
		},
		Unique: true,
	},
	{
		Field: []repos.IndexKey{
			{Key: "metadata.name", Value: repos.IndexAsc},
			{Key: "clusterName", Value: repos.IndexAsc},
			{Key: "accountName", Value: repos.IndexAsc},
		},
		Unique: true,
	},
}
