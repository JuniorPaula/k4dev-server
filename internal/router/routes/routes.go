package routes

import (
	"knowledge-api/internal/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI     string
	Method  string
	Func    func(w http.ResponseWriter, r *http.Request)
	HasAuth bool
}

func SetUpRoutes(r *mux.Router) *mux.Router {
	routes := homeRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, usersRoutes...)

	for _, route := range routes {

		if route.HasAuth {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticated(route.Func))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Func)).Methods(route.Method)
		}

	}

	return r
}
