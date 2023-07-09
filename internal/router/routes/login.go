package routes

import (
	"knowledge-api/internal/handlers"
	"net/http"
)

var loginRoute = Route{
	URI:     "/login",
	Method:  http.MethodPost,
	Func:    handlers.Login,
	HasAuth: false,
}
