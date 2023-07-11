package handlers

import (
	"encoding/json"
	"io"
	"knowledge-api/internal/models"
	"knowledge-api/internal/usecases"
	"knowledge-api/internal/utils"
	"net/http"
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
