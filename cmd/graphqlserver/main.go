package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/shawnzxx/graphql-twitter/config"
	"github.com/shawnzxx/graphql-twitter/graph"
	"github.com/shawnzxx/graphql-twitter/postgres"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()

	config.LoadEnv(".env")

	conf := config.New()

	db := postgres.New(ctx, conf)

	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RedirectSlashes)
	router.Use(middleware.Timeout(time.Second * 60))

	router.Handle("/", playground.Handler("Twitter clone", "/query"))
	router.Handle("/query", handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{},
			},
		),
	))

	log.Println("server running")
	log.Fatal(http.ListenAndServe(":8080", router))
}
