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

func FindAllUsersUSecase() ([]models.User, error) {
	db, err := database.Connect_MySQL()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	repo := repository.NewUsersRepository(db)

	users, err := repo.FindAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func FindUserByIDUSecase(id int64) (models.User, error) {
	db, err := database.Connect_MySQL()
	if err != nil {
		return models.User{}, err
	}
	defer db.Close()

	repo := repository.NewUsersRepository(db)

	user, err := repo.FindUserByID(id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func UpdateUserUSecase(userID int64, u models.User) error {

	if err := u.HanlderUser("updated"); err != nil {
		return err
	}

	db, err := database.Connect_MySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	repo := repository.NewUsersRepository(db)
	if err := repo.UpdateUser(userID, u); err != nil {
		return err
	}

	return nil
}
