package routes

import (
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
	routes = append(routes, usersRoutes...)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return r
}
