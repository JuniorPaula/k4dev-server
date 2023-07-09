package repository

import (
	"database/sql"
	"knowledge-api/internal/models"
)

type users struct {
	DB *sql.DB
}

func NewUsersRepository(db *sql.DB) *users {
	return &users{
		DB: db,
	}
}

func (u *users) CreateUser(user models.User) (int64, error) {
	statment, err := u.DB.Prepare("insert into users (name, email, password, admin) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statment.Close()

	result, err := statment.Exec(user.Name, user.Email, user.Password, user.Admin)
	if err != nil {
		return 0, err
	}

	lstID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int64(lstID), nil
}

func (u *users) FindAllUsers() ([]models.User, error) {
	rows, err := u.DB.Query("select id, name, email, password, admin from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *users) FindUserByID(id int64) (models.User, error) {
	rows, err := u.DB.Query("select id, name, email, password, admin from users where id = ?", id)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (u *users) UpdateUser(id int64, user models.User) error {
	statment, err := u.DB.Prepare("update users set name = ?, email = ? where id = ?")
	if err != nil {
		return err
	}
	defer statment.Close()

	if _, err := statment.Exec(user.Name, user.Email, id); err != nil {
		return err
	}

	return nil
}

func (u *users) FindUserByEmail(email string) (models.User, error) {
	rows, err := u.DB.Query("select id, name, email, password, admin from users where email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (u *users) DeleteUser(userID int64) error {
	statment, err := u.DB.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}
	defer statment.Close()

	if _, err := statment.Exec(userID); err != nil {
		return err
	}

	return nil
}
