package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/traineira/stretch/cmd/graphql/graph"
	"github.com/traineira/stretch/cmd/graphql/graph/generated"
	"github.com/traineira/stretch/pkg/postgres"
)

const defaultPort = "8080"

var port = os.Getenv("PORT")
var dbUserName = os.Getenv("POSTGRES_USER")
var dbPassword = os.Getenv("POSTGRES_PASSWORD")
var dbURL = os.Getenv("POSTGRES_URL")
var dbName = os.Getenv("POSTGRES_DB")

func main() {
	if port == "" {
		port = defaultPort
	}

	toDoService := &postgres.ToDoService{
		DbUserName: dbUserName,
		DbPassword: dbPassword,
		DbURL:      dbURL,
		DbName:     dbName,
	}

	err := toDoService.Initialise()
	if err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					ToDo: toDoService,
				}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
