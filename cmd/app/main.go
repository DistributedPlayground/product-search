package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	env "github.com/DistributedPlayground/go-lib/config"
	"github.com/DistributedPlayground/product-search/graph"
	gql_api "github.com/DistributedPlayground/product-search/graph/api"
	"github.com/DistributedPlayground/product-search/pkg/store"
	"github.com/DistributedPlayground/products/config"
)

func main() {
	err := env.LoadEnv(&config.Var)
	if err != nil {
		panic(err)
	}

	db := store.MustNewMongo()
	srv := handler.NewDefaultServer(gql_api.NewExecutableSchema(gql_api.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Var.PORT)
	log.Fatal(http.ListenAndServe(":"+config.Var.PORT, nil))

}
