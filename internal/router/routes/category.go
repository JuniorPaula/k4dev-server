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
	{
		URI:     "/category/{categoryId}",
		Method:  http.MethodPut,
		Func:    handlers.UpdateCategory,
		HasAuth: true,
	},
	{
		URI:     "/category/delete/{categoryId}",
		Method:  http.MethodDelete,
		Func:    handlers.DeleteCategory,
		HasAuth: true,
	}, {
		URI:     "/category",
		Method:  http.MethodGet,
		Func:    handlers.FindAllCategories,
		HasAuth: true,
	},
	{
		URI:     "/category/tree",
		Method:  http.MethodGet,
		Func:    handlers.FindCategoryWithTree,
		HasAuth: true,
	},
	{
		URI:     "/category/{categoryId}",
		Method:  http.MethodGet,
		Func:    handlers.FindCategoryByID,
		HasAuth: true,
	},
}
