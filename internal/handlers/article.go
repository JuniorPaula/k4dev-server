package handlers

import (
	"encoding/json"
	"io"
	"knowledge-api/internal/auth"
	"knowledge-api/internal/models"
	"knowledge-api/internal/usecases"
	"knowledge-api/internal/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func FindArticleByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	articleID, err := strconv.ParseInt(params["articleId"], 10, 64)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	article, err := usecases.FindArticleByIDUsecase(articleID)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, article)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	articleID, err := strconv.ParseInt(params["articleId"], 10, 64)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	userIDInToken, err := auth.GetUserID(r)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

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

	if err = usecases.UpdateArticlesUsecase(articleID, userIDInToken, article); err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	articleID, err := strconv.ParseInt(params["articleId"], 10, 64)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	userIDInToken, err := auth.GetUserID(r)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	if err = usecases.DeleteArticleUsecase(articleID, userIDInToken); err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)
}
