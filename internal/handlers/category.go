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

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	var category models.Category
	if err = json.Unmarshal(body, &category); err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	category, err = usecases.CreateCategoryUSecase(category)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, category)
}

func FindAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := usecases.FindAllCategoriesUsecase()
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}
	var category models.Category
	categories = category.WithPath(categories)

	utils.WriteJSON(w, http.StatusOK, categories)
}

func FindCategoryWithTree(w http.ResponseWriter, r *http.Request) {
	categories, err := usecases.FindAllCategoriesUsecase()
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	var category models.Category
	categories = category.WithPath(categories)
	categories = category.ToTree(categories, nil)

	utils.WriteJSON(w, http.StatusOK, categories)
}

func FindCategoryByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	categoryID, err := strconv.ParseInt(params["categoryId"], 10, 64)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	category, err := usecases.FindCategoryByIDUsecase(categoryID)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, category)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	categoryID, err := strconv.ParseInt(params["categoryId"], 10, 64)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	userIDInToken, err := auth.GetUserID(r)
	if err != nil {
		utils.ErrorJSON(w, http.StatusUnauthorized, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	var category models.Category
	if err := json.Unmarshal(body, &category); err != nil {
		utils.ErrorJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = usecases.UpdateCategoryUsecase(categoryID, userIDInToken, category)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	categoryID, err := strconv.ParseInt(params["categoryId"], 10, 64)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	userIDInToken, err := auth.GetUserID(r)
	if err != nil {
		utils.ErrorJSON(w, http.StatusUnauthorized, err)
		return
	}

	var category models.Category

	err = usecases.DeleteCategoryUsecase(categoryID, userIDInToken, category)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)
}
