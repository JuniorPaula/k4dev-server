package routes

import (
	"knowledge-api/internal/handlers"
	"net/http"
)

var loginRoute = []Route{
	{
		URI:     "/login",
		Method:  http.MethodPost,
		Func:    handlers.Login,
		HasAuth: false,
	},
	{
		URI:     "/signup",
		Method:  http.MethodPost,
		Func:    handlers.Signup,
		HasAuth: false,
	},
	{
		URI:     "/validateToken",
		Method:  http.MethodPost,
		Func:    handlers.VerifyIfTokenIsValid,
		HasAuth: false,
	},
}
