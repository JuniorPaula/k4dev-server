package routes

import (
	"knowledge-api/internal/handlers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:     "/users",
		Method:  http.MethodPost,
		Func:    handlers.CreateUser,
		HasAuth: false,
	},
}
