package list

import (
	"errors"
	"fmt"

	"github.com/kloudlite/kl/domain/client"
	"github.com/kloudlite/kl/domain/server"
	fn "github.com/kloudlite/kl/pkg/functions"
	"github.com/kloudlite/kl/pkg/ui/table"
	"github.com/kloudlite/kl/pkg/ui/text"

	"github.com/spf13/cobra"
)

var clustersCmd = &cobra.Command{
	Use:   "clusters",
	Short: "list all the clusters accessible to you",
	Long: `List Clusters

This command will provide the list of all the clusters that's accessible to you. 

Examples:
  # list clusters accessible to you
  kl list clusters

Note: selected project will be highlighted with green color.

`,
	Run: func(cmd *cobra.Command, _ []string) {
		err := listClusters(cmd)
		if err != nil {
			fn.PrintError(err)
			return
		}
	},
}

func listClusters(cmd *cobra.Command) error {
	clusters, err := server.ListClusters()
	if err != nil {
		return err
	}

	if len(clusters) == 0 {
		return errors.New("no clusters found")
	}

	clusterName, _ := client.CurrentClusterName()

	header := table.Row{table.HeaderText("name"), table.HeaderText("id"), table.HeaderText("ready")}
	rows := make([]table.Row, 0)

	for _, a := range clusters {
		rows = append(rows, table.Row{
			func() string {
				if a.Metadata.Name == clusterName {
					return text.Colored(fmt.Sprint("*", a.DisplayName), 2)
				}
				return a.DisplayName
			}(),
			func() string {
				if a.Metadata.Name == clusterName {
					return text.Colored(a.Metadata.Name, 2)
				}
				return a.Metadata.Name
			}(),

			func() string {
				if a.Metadata.Name == clusterName {
					return text.Colored(a.Status.IsReady, 2)
				}
				return fmt.Sprint(a.Status.IsReady)
			}(),
		})
	}

	fmt.Println(table.Table(&header, rows, cmd))

	if s := fn.ParseStringFlag(cmd, "output"); s == "table" {
		table.TotalResults(len(clusters), true)
	}

	return nil
}

func init() {
	clustersCmd.Aliases = append(clustersCmd.Aliases, "cluster")
	clustersCmd.Aliases = append(clustersCmd.Aliases, "clus")
	fn.WithOutputVariant(clustersCmd)
}
