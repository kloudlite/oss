package app

import (
	"kloudlite.io/pkg/messaging"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"kloudlite.io/apps/console/internal/app/graph"
	"kloudlite.io/apps/console/internal/app/graph/generated"
	"kloudlite.io/apps/console/internal/domain"
	"kloudlite.io/apps/console/internal/domain/entities"
	"kloudlite.io/pkg/repos"
)

var Module = fx.Module(
	"app",
	// Create Repos
	fx.Provide(func(db *mongo.Database) (
		repos.DbRepo[*entities.Cluster],
		repos.DbRepo[*entities.Device],
	) {
		deviceRepo := repos.NewMongoRepoAdapter[*entities.Device](db, "devices", "dev")
		clusterRepo := repos.NewMongoRepoAdapter[*entities.Cluster](db, "clusters", "cluster")
		return clusterRepo, deviceRepo
	}),
	fx.Provide(func(messagingCli *messaging.KafkaClient) (messaging.Producer[messaging.Json], error) {
		return messaging.NewKafkaProducer[messaging.Json](messagingCli)
	}),
	// Load Domain
	domain.Module,
	// Create GQL Handler from domain
	fx.Provide(func(d domain.Domain) http.Handler {
		server := http.NewServeMux()
		server.HandleFunc("/", playground.Handler("Graphql playground", "/query"))

		gqlServer := handler.NewDefaultServer(
			generated.NewExecutableSchema(
				generated.Config{Resolvers: &graph.Resolver{Domain: d}},
			),
		)

		server.Handle("/query", gqlServer)
		c := cors.New(cors.Options{
			AllowedOrigins:   []string{"http://localhost:4001", "https://studio.apollographql.com"},
			AllowCredentials: true,
			AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		})

		return c.Handler(server)
	}),
)
