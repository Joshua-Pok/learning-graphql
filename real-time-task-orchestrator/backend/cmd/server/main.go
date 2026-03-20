package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Joshua-Pok/task-orchestrator/internal/gql"
	"github.com/Joshua-Pok/task-orchestrator/internal/gql/generated"
	"github.com/Joshua-Pok/task-orchestrator/internal/task"
)

func main() {
	store := &task.Store{}
	resolver := &gql.Resolver{
		Store: store,
	}

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{Resolvers: resolver}),
	)

	http.Handle("/query", srv)
	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))

	log.Println("Server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
