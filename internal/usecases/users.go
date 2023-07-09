package usecases

import (
	"errors"
	"knowledge-api/internal/auth"
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

func UpdateUserUSecase(userID, userIDInToken int64, u models.User) error {
	if err := u.HanlderUser("updated"); err != nil {
		return err

	}

	db, err := database.Connect_MySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	repo := repository.NewUsersRepository(db)
	userFromDB, err := repo.FindUserByID(userIDInToken)
	if err != nil {
		return err
	}

	if userIDInToken != userID && !userFromDB.Admin {
		return errors.New("you don't have permission to update this user")
	}

	if u.Name == "" {
		u.Name = userFromDB.Name
	}

	if u.Email == "" {
		u.Email = userFromDB.Email
	}

	if err := repo.UpdateUser(userID, u); err != nil {
		return err
	}

	return nil
}

func DeleteUserUsecase(userID, userIDInToken int64) error {
	db, err := database.Connect_MySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	repo := repository.NewUsersRepository(db)
	userFromDB, err := repo.FindUserByID(userIDInToken)
	if err != nil {
		return err
	}

	if userIDInToken != userID && !userFromDB.Admin {
		return errors.New("you don't have permission to delete this user")
	}

	if err := repo.DeleteUser(userID); err != nil {
		return err
	}

	return nil
}

func UpdatedUserPasswordUsecase(userID, userIDInToken int64, password models.PasswordDTO) error {
	db, err := database.Connect_MySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	repo := repository.NewUsersRepository(db)
	userFromDB, err := repo.FindUserByID(userIDInToken)
	if err != nil {
		return err
	}

	if userIDInToken != userID && !userFromDB.Admin {
		return errors.New("you dont't have permission of the operation")
	}

	passwordFromDB, err := repo.FindPasswordByUserID(userID)
	if err != nil {
		return err
	}

	if err = auth.CompareHash(passwordFromDB, password.OldPassword); err != nil {
		return errors.New("passwords does not match")
	}

	passwordWithHash, err := auth.Hash(password.NewPassword)
	if err != nil {
		return err
	}

	if err = repo.UpdatePassword(userID, string(passwordWithHash)); err != nil {
		return err
	}

	return nil
}
