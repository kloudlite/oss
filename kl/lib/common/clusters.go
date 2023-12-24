package common

import (
	"github.com/kloudlite/kl/lib/server"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/pkg/errors"
)

func SelectCluster(args []string) (string, error) {
	clusterName := ""
	if len(args) >= 1 {
		clusterName = args[0]
	}
	clusters, err := server.GetClusters()
	if err != nil {
		return "", err
	}

	if clusterName != "" {
		for _, a := range clusters {
			if a.Metadata.Name == clusterName {
				return a.Metadata.Name, nil
			}
		}
		return "", errors.New("you don't have access to this cluster")
	}

	selectedIndex, err := fuzzyfinder.Find(
		clusters,
		func(i int) string {
			return clusters[i].DisplayName
		},
		fuzzyfinder.WithPromptString("Use Cluster >"),
	)

	if err != nil {
		return "", err
	}

	return clusters[selectedIndex].Metadata.Name, nil
}
