package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bamboooo-dev/meshi-api/graph"
	"github.com/bamboooo-dev/meshi-api/graph/generated"

	"go.uber.org/zap"
)

var revision = "undefined"

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panic '%v' captured\n", err)
		}
	}()

	fmt.Printf("Version is %s\n", revision)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, _ := zap.NewProduction()
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Error(err.Error())
			return
		}
	}()

	sugar := logger.Sugar()

	resolver := graph.NewResolver(sugar)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Print("connect to http://localhost:8080/ for GraphQL playground")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		sugar.Error(ctx, err)
	}
}
