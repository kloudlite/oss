package framework

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"kloudlite.io/apps/wireguard/internal/app"
	"kloudlite.io/pkg/config"
	gql_server "kloudlite.io/pkg/gql-server"
	"kloudlite.io/pkg/logger"
	mongo_db "kloudlite.io/pkg/mongo-db"
	"net/http"
)

type Env struct {
	MongoUri    string `env:"MONGO_URI", required:"true"`
	MongoDbName string `env:"MONGO_DB_NAME", required:"true"`
	Port        uint32 `env:"PORT", required:"true"`
	IsDev       bool   `env:"DEV", default:"false"`
}

var Module = fx.Module("framework",
	// Load Config from Env
	fx.Provide(func() (*Env, error) {
		var envC Env
		err := config.LoadConfigFromEnv(&envC)
		if err != nil {
			fmt.Println(err, "failed to load env asdkadhaskda dkjahd kashda kh")
			return nil, fmt.Errorf("not able to load ENV: %w", err)
		}
		return &envC, err
	}),
	// Setup Logger
	fx.Provide(func(env *Env) logger.Logger {
		return logger.NewLogger(env.IsDev)
	}),
	// Create DB Client
	fx.Provide(func(env *Env) (*mongo.Database, error) {
		return mongo_db.NewMongoDatabase(env.MongoUri, env.MongoDbName)
	}),
	// Connect DB Client
	fx.Invoke(func(lifecycle fx.Lifecycle, db *mongo.Database) {
		lifecycle.Append(fx.Hook{
			OnStart: func(context.Context) error {
				return db.Client().Connect(context.Background())
			},
		})
	}),
	// Load App Module
	app.Module,
	// start http server
	fx.Invoke(func(lf fx.Lifecycle, env *Env, logger logger.Logger, gqlHandler http.Handler) {
		lf.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				return gql_server.StartGQLServer(ctx, env.Port, gqlHandler, logger)
			},
			OnStop: func(context.Context) error {
				return nil
			},
		})
	}),
)
