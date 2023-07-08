package routes

import (
	"knowledge-api/internal/handlers"
	"net/http"
)

var homeRoutes = []Route{
	{
		URI:     "/",
		Method:  http.MethodGet,
		Func:    handlers.Home,
		HasAuth: false,
	},
}
