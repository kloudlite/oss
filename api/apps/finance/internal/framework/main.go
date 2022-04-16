package framework

import (
	"go.uber.org/fx"
	"kloudlite.io/pkg/config"
	"kloudlite.io/pkg/repos"
)

type Env struct {
	DBName string `env:"DB_NAME"`
	DBUrl  string `env:"DB_URL"`
}

func (e *Env) GetMongoConfig() (url string, dbName string) {
	return e.DBUrl, e.DBName
}

var Module = fx.Module("framework",
	fx.Provide(config.LoadEnv[*Env]()),
	repos.NewMongoClientFx[*Env](),
)
