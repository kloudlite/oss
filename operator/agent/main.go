package agent

import (
	"bytes"
	"encoding/json"
	"operators.kloudlite.io/lib/errors"
	"operators.kloudlite.io/lib/logging"
	"operators.kloudlite.io/lib/redpanda"
	"os"
	"os/exec"
	"sigs.k8s.io/yaml"
)

func Run(c *redpanda.Consumer, logger logging.Logger) {
	c.StartConsuming(
		func(m *redpanda.Message) error {
			logger.Debugf("action=%s, payload=%s\n", m.Action, m.Payload)
			c := exec.Command("kubectl", m.Action, "-f", "-")
			jb, err := json.Marshal(m.Payload)
			if err != nil {
				return errors.NewEf(err, "could not unmarshal into []byte")
			}
			yb, err := yaml.JSONToYAML(jb)
			if err != nil {
				return errors.NewEf(err, "could not convert JSON to YAML")
			}

			c.Stdin = bytes.NewBuffer(yb)
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			return c.Run()
		},
	)
}
