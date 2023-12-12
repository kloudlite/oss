package framework

import (
	"context"
	"fmt"
	"github.com/kloudlite/api/apps/webhooks/internal/app"
	"github.com/kloudlite/api/apps/webhooks/internal/env"
	httpServer "github.com/kloudlite/api/pkg/http-server"
	"github.com/kloudlite/api/pkg/logging"
	"github.com/kloudlite/api/pkg/redpanda"
	"go.uber.org/fx"
)

type fm struct {
	*env.Env
}

func (f fm) GetKafkaSASLAuth() *redpanda.KafkaSASLAuth {
	return nil
	// return &redpanda.KafkaSASLAuth{
	// 	SASLMechanism: redpanda.ScramSHA256,
	// 	User:          v.KafkaUsername,
	// 	Password:      v.KafkaPassword,
	// }
}

func (f fm) GetBrokers() string {
	return f.KafkaBrokers
}

func (f fm) GetHttpPort() uint16 {
	return f.HttpPort
}

func (f fm) GetHttpCors() string {
	return ""
}

var Module = fx.Module(
	"framework",
	fx.Provide(
		func(vars *env.Env) *fm {
			return &fm{Env: vars}
		},
	),

	redpanda.NewClientFx[*fm](),
	fx.Provide(func(logger logging.Logger) httpServer.Server {
		corsOrigins := "https://studio.apollographql.com"
		return httpServer.NewServer(httpServer.ServerArgs{Logger: logger, CorsAllowOrigins: &corsOrigins})
	}),

	fx.Invoke(func(lf fx.Lifecycle, server httpServer.Server, ev *env.Env) {
		lf.Append(fx.Hook{
			OnStart: func(context.Context) error {
				return server.Listen(fmt.Sprintf(":%d", ev.HttpPort))
			},
			OnStop: func(context.Context) error {
				return server.Close()
			},
		})
	}),
	app.Module,
)
