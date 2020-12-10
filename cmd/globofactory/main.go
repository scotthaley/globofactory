package main

import (
	"github.com/graphql-go/handler"
	"github.com/scotthaley/globofactory/internal/api"
	"github.com/scotthaley/globofactory/internal/database"
	"net/http"
)

func main() {
	database.InitDB()

	h := handler.New(&handler.Config{
		Schema: &api.Schema,
		Pretty: true,
	})

	// static file server for GraphQL playground
	fs := http.FileServer(http.Dir("static"))

	http.Handle("/graphql", h)
	http.Handle("/", fs)
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		panic(err)
	}
}