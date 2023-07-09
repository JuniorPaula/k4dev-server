package usecases

import (
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
		return models.AuthDTO{}, err
	}

	token, err := auth.TokenGenerator(userFromDB.ID)
	if err != nil {
		return models.AuthDTO{}, err
	}

	userID := strconv.FormatInt(userFromDB.ID, 10)

	return models.AuthDTO{
		UserID: userID,
		Token:  token,
	}, nil
}
