package routes

import (
	"knowledge-api/internal/handlers"
	"net/http"
)

var articleRoutes = []Route{
	{
		URI:     "/articles",
		Method:  http.MethodPost,
		Func:    handlers.CreateArticle,
		HasAuth: true,
	},
}