package usecases

import (
	"knowledge-api/internal/database"
	"knowledge-api/internal/models"
	"knowledge-api/internal/repository"
)

func CreateUserUSecase(u models.User) (models.User, error) {
	if err := u.HanlderUser("create"); err != nil {
		return models.User{}, err
	}

	db, err := database.Connect_MySQL()
	if err != nil {
		return models.User{}, err
	}
	defer db.Close()

	repo := repository.NewUsersRepository(db)

	u.ID, err = repo.CreateUser(u)
	if err != nil {
		return models.User{}, err
	}

	return u, nil
}
