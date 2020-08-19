package main

import (
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type RootResolver struct {
}

func (_ *RootResolver) Hello() string {
	return "Hello, @103cuong"
}

func main() {
	s := `
		type Query {
		hello: String!
	}
	`
	schema := graphql.MustParseSchema(s, &RootResolver{})
	http.Handle("/graphql", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
