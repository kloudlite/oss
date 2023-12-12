package framework

import (
	"github.com/kloudlite/api/apps/worker-audit-logging/internal/app"
	"github.com/kloudlite/api/apps/worker-audit-logging/internal/env"
	"github.com/kloudlite/api/pkg/redpanda"
	repos "github.com/kloudlite/api/pkg/repos"
	"go.uber.org/fx"
	"strings"
)

type redpandaCfg struct {
	ev *env.Env
}

func (r redpandaCfg) GetSubscriptionTopics() []string {
	return strings.Split(r.ev.KafkaSubscriptionTopics, ",")
}

func (r redpandaCfg) GetConsumerGroupId() string {
	return r.ev.KafkaConsumerGroupId
}

func (r redpandaCfg) GetBrokers() (brokers string) {
	return r.ev.KafkaBrokers
}

func (r redpandaCfg) GetKafkaSASLAuth() *redpanda.KafkaSASLAuth {
	return &redpanda.KafkaSASLAuth{
		SASLMechanism: redpanda.ScramSHA256,
		User:          r.ev.KafkaUsername,
		Password:      r.ev.KafkaPassword,
	}
}

type eventsDbCfg struct {
	ev *env.Env
}

func (db eventsDbCfg) GetMongoConfig() (url string, dbName string) {
	return db.ev.EventsDbUri, db.ev.EventsDbName
}

var Module fx.Option = fx.Module("framework",
	fx.Provide(func(ev *env.Env) *redpandaCfg {
		return &redpandaCfg{ev: ev}
	}),
	redpanda.NewClientFx[*redpandaCfg](),
	redpanda.NewConsumerFx[*redpandaCfg](),
	redpanda.NewProducerFx[redpanda.Client](),

	fx.Provide(func(ev *env.Env) *eventsDbCfg {
		return &eventsDbCfg{ev: ev}
	}),
	repos.NewMongoClientFx[*eventsDbCfg](),
	app.Module,
)
