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
	{
		URI:     "/articles/{articleId}",
		Method:  http.MethodPut,
		Func:    handlers.UpdateArticle,
		HasAuth: true,
	},
	{
		URI:     "/articles/{articleId}",
		Method:  http.MethodDelete,
		Func:    handlers.DeleteArticle,
		HasAuth: true,
	},
	{
		URI:     "/articles/{articleId}",
		Method:  http.MethodGet,
		Func:    handlers.FindArticleByID,
		HasAuth: true,
	},
	{
		URI:     "/articles",
		Method:  http.MethodGet,
		Func:    handlers.FindAllArticles,
		HasAuth: true,
	},
	{
		URI:     "/category/{categoryId}/articles",
		Method:  http.MethodGet,
		Func:    handlers.FindCategoryWithChildren,
		HasAuth: true,
	},
}
