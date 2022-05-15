package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/shawnzxx/graphql-twitter/config"
	"github.com/shawnzxx/graphql-twitter/domain"
	"github.com/shawnzxx/graphql-twitter/graph"
	"github.com/shawnzxx/graphql-twitter/postgres"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()

	// load configuration
	config.LoadEnv(".env")
	conf := config.New()

	// init db pool and configure
	db := postgres.New(ctx, conf)

	// migrate db when every time app restart
	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()

	// middleware section
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RedirectSlashes)
	router.Use(middleware.Timeout(time.Second * 60))

	// 1: inject db into Repos layer
	userRepo := postgres.NewUserRepo(db)
	// 2: inject db repo into Domain Service layer
	authService := domain.NewAuthService(userRepo)

	// graphql settings
	router.Handle("/", playground.Handler("Twitter clone", "/query"))
	router.Handle("/query", handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				// 3: inject domain service into graphql handler
				Resolvers: &graph.Resolver{
					AuthService: authService,
				},
			},
		),
	))

	// start server
	log.Fatal(http.ListenAndServe(":8080", router))
}
