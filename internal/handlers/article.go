package handlers

import (
	"encoding/json"
	"io"
	"knowledge-api/internal/models"
	"knowledge-api/internal/usecases"
	"knowledge-api/internal/utils"
	"net/http"
)

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	var article models.Article
	if err = json.Unmarshal(body, &article); err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	article, err = usecases.CreateArticleUsecase(article)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, article)
}
