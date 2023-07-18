package handlers

import (
	"knowledge-api/internal/usecases"
	"knowledge-api/internal/utils"
	"net/http"
)

func FindStat(w http.ResponseWriter, r *http.Request) {
	stats, _ := usecases.FindStatUsecase()

	utils.WriteJSON(w, http.StatusOK, stats)
}
