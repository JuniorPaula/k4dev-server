package handlers

import (
	"knowledge-api/internal/usecases"
	"knowledge-api/internal/utils"
	"net/http"
)

func FindStat(w http.ResponseWriter, r *http.Request) {
	stats, err := usecases.FindStatUsecase()
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, stats)
}
