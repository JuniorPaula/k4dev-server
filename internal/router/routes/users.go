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
	{
		URI:     "/users",
		Method:  http.MethodGet,
		Func:    handlers.FindAllUsers,
		HasAuth: true,
	},
	{
		URI:     "/users/{userId}",
		Method:  http.MethodGet,
		Func:    handlers.FindUserByID,
		HasAuth: true,
	},
	{
		URI:     "/users/{userId}",
		Method:  http.MethodPut,
		Func:    handlers.UpdateUser,
		HasAuth: true,
	},
	{
		URI:     "/users/{userId}",
		Method:  http.MethodDelete,
		Func:    handlers.DeleteUser,
		HasAuth: true,
	},
	{
		URI:     "/users/{userId}/update-password",
		Method:  http.MethodPost,
		Func:    handlers.UpdatePassword,
		HasAuth: true,
	},
	{
		URI:     "/users/{userId}/update-role",
		Method:  http.MethodPost,
		Func:    handlers.UpdateUserRole,
		HasAuth: true,
	},
}
