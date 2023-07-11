package routes

import (
	"knowledge-api/internal/handlers"
	"net/http"
)

var categoryRoutes = []Route{
	{
		URI:     "/category",
		Method:  http.MethodPost,
		Func:    handlers.CreateCategory,
		HasAuth: true,
	},
}
