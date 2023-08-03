package router

import (
	"knowledge-api/internal/config"
	"knowledge-api/internal/router/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func HanlderRoutes() *mux.Router {
	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{config.FrontendURL},
		AllowedMethods: []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	r.Use(c.Handler)

	return routes.SetUpRoutes(r)
}
