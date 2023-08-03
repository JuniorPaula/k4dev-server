package router

import (
	"knowledge-api/internal/router/routes"

	"github.com/gorilla/mux"
)

func HanlderRoutes() *mux.Router {
	r := mux.NewRouter()

	return routes.SetUpRoutes(r)
}
