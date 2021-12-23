package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/graphql-go/graphql"
	"github.com/traineira/stretch/gql"
	"github.com/traineira/stretch/postgres"
	"github.com/traineira/stretch/server"
)

func main() {
	router, db := initializeAPI()
	defer db.Close()

	log.Fatal(http.ListenAndServe(":4000", router))
}

func initializeAPI() (*chi.Mux, *postgres.Db) {
	router := chi.NewRouter()

	db, err := postgres.New(
		postgres.ConnString("localhost", 5432, "joaocarvalho", "stretch_db"),
	)

	if err != nil {
		log.Fatal(err)
	}

	rootQuery := gql.NewRoot(db)

	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)

	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	s := server.Server{
		GqlSchema: &sc,
	}
}
