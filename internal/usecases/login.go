package usecases

import (
	"errors"
	"knowledge-api/internal/auth"
	"knowledge-api/internal/database"
	"knowledge-api/internal/models"
	"knowledge-api/internal/repository"
	"strconv"
)

func LoginUsecase(user models.User) (models.AuthDTO, error) {
	db, err := database.Connect_MySQL()
	if err != nil {
		return models.AuthDTO{}, err
	}
	defer db.Close()

	repo := repository.NewUsersRepository(db)
	userFromDB, err := repo.FindUserByEmail(user.Email)
	if err != nil {
		return models.AuthDTO{}, err
	}

	if err = auth.CompareHash(userFromDB.Password, user.Password); err != nil {
		return models.AuthDTO{}, errors.New("invalid credentials")
	}

	token, err := auth.TokenGenerator(userFromDB.ID)
	if err != nil {
		return models.AuthDTO{}, err
	}

	userID := strconv.FormatInt(userFromDB.ID, 10)

	return models.AuthDTO{
		UserID: userID,
		Token:  token,
		Role:   userFromDB.Admin,
	}, nil
}

func SignupUsecase(user models.User) (string, error) {
	if err := user.HanlderUser("create"); err != nil {
		return "", err
	}

	db, err := database.Connect_MySQL()
	if err != nil {
		return "", err
	}
	defer db.Close()

	repo := repository.NewUsersRepository(db)
	userFromDB, err := repo.FindUserByEmail(user.Email)
	if err != nil {
		return "", err
	}

	if userFromDB.ID != 0 {
		return "", errors.New("user already exists")
	}

	_, err = repo.CreateUser(user)
	if err != nil {
		return "", err
	}

	return "user created on success", nil
}
