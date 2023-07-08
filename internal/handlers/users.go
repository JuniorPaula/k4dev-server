package handlers

import (
	"encoding/json"
	"io"
	"knowledge-api/internal/database"
	"knowledge-api/internal/models"
	"knowledge-api/internal/repository"
	"knowledge-api/internal/utils"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(reqBody, &user); err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	if err = user.HanlderUser("create"); err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConnectToDB()
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewUsersRepository(db)

	user.ID, err = repo.CreateUser(user)
	if err != nil {
		utils.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, user)
}
