package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"workspace-service/core/config"
	"workspace-service/core/database"
	"workspace-service/core/database/models"
	"workspace-service/core/vault"
	"workspace-service/graphql"
	resolvers "workspace-service/graphql/resolvers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"

	graphql_client "github.com/hasura/go-graphql-client"
	"golang.org/x/oauth2"
)

const defaultPort = "8080"

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	microServicesConfig := config.GetToucanMicroServicesConfig()
	db, err := database.Connect(config.GetDatabaseConfig())
	if err != nil {
		panic(err)
	}
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	db.AutoMigrate(&models.User{}, &models.Group{}, &models.App{})

	httpClient := oauth2.NewClient(context.Background(), vault.GetTokenSource(microServicesConfig.Vault))
	curityClient := graphql_client.NewClient(microServicesConfig.Curity.Uri, httpClient)

	resolvers := resolvers.NewResolvers(db, curityClient)
	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: resolvers}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
