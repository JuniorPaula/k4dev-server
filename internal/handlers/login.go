package handlers

import (
	"encoding/json"
	"io"
	"knowledge-api/internal/auth"
	"knowledge-api/internal/models"
	"knowledge-api/internal/usecases"
	"knowledge-api/internal/utils"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(reqBody, &user); err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	authDTO, err := usecases.LoginUsecase(user)
	if err != nil {
		utils.ErrorJSON(w, http.StatusUnauthorized, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, authDTO)

}

func Signup(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(reqBody, &user); err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	user.Admin = false

	authDTO, err := usecases.SignupUsecase(user)
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, authDTO)
}

func VerifyIfTokenIsValid(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	var authData models.AuthDTO
	if err = json.Unmarshal(reqBody, &authData); err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	isValid := auth.CheckToken(authData.Token)

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"isValid": isValid,
	})
}
