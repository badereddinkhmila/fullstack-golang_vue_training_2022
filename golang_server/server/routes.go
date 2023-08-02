package server

import (
	"fmt"
	"time"

	"com.fullstack.ecommerce/server/auth"
	"com.fullstack.ecommerce/utils/config"

	gqlResolver "com.fullstack.ecommerce/server/gql"
	"com.fullstack.ecommerce/server/gql/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	chi "github.com/go-chi/chi/v5"
	middleware "github.com/go-chi/chi/v5/middleware"
	cors "github.com/go-chi/cors"
	httprate "github.com/go-chi/httprate"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	appConfig := config.App()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(httprate.LimitByIP(1000, 2*time.Second))
	fmt.Println(appConfig.Cors)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   appConfig.Cors,
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))
	router.Use(middleware.Logger)
	//JWT Auth Middleware
	router.Use(auth.AuthenticationMiddleware())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: gqlResolver.Resolver()}))
	router.Route("/", func(router chi.Router) {
		router.Handle("/playground", playground.Handler("Estale GQL Playground", "/query"))
		router.Handle("/query", srv)
	})

	return router
}
