package routes

import (
	"knowledge-api/internal/handlers"
	"net/http"
)

var statsRoutes = []Route{
	{
		URI:     "/stats",
		Method:  http.MethodGet,
		Func:    handlers.FindStat,
		HasAuth: true,
	},
}
