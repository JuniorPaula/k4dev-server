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

	user, err = usecases.CreateUserUSecase(user)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, user)
}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := usecases.FindAllUsersUSecase()
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, users)
}

func FindUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseInt(params["userId"], 10, 64)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	user, err := usecases.FindUserByIDUSecase(userID)
	if err != nil {
		utils.ErrorJSON(w, http.StatusNotFound, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseInt(params["userId"], 10, 60)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	userIDInToken, err := auth.GetUserID(r)
	if err != nil {
		utils.ErrorJSON(w, http.StatusUnauthorized, err)
		return
	}

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

	err = usecases.UpdateUserUSecase(userID, userIDInToken, user)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseInt(params["userId"], 10, 64)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	userIDInToken, err := auth.GetUserID(r)
	if err != nil {
		utils.ErrorJSON(w, http.StatusUnauthorized, err)
		return
	}

	err = usecases.DeleteUserUsecase(userID, userIDInToken)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)
}
