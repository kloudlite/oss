package boxpkg

import (
	"fmt"

	"github.com/kloudlite/kl/pkg/ui/table"
	"github.com/kloudlite/kl/pkg/ui/text"
)

func trimePref(s string, length int) string {

	if len(s) < length {
		return s
	}

	return fmt.Sprintf("...%s", s[len(s)-length:])
}

func (c *client) ListBox() error {
	conts, err := c.listContainer(map[string]string{
		CONT_MARK_KEY: "true",
	})
	if err != nil && err != notFoundErr {
		return err
	}

	if err == notFoundErr {
		return fmt.Errorf("no running container found")
	}

	header := table.Row{table.HeaderText("container name"), table.HeaderText("path")}
	rows := make([]table.Row, 0)

	for _, a := range conts {
		rows = append(rows, table.Row{
			func() string {
				if a.Name == c.containerName {
					return text.Colored(a.Name, 2)
				}
				return a.Name
			}(),
			func() string {
				pth := trimePref(a.Labels[CONT_PATH_KEY], 50)

				if a.Name == c.containerName {
					return text.Colored(pth, 2)
				}
				return pth
			}(),
		})
	}

	fmt.Println(table.Table(&header, rows, c.cmd))

	table.TotalResults(len(conts), true)

	return nil
}
