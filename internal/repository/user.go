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
	statment, err := u.DB.Prepare("insert into users (name, email, password) values ($1, $2, $3) returning id")
	if err != nil {
		return 0, nil
	}
	defer statment.Close()

	result, err := statment.Exec(user.Name, user.Email, user.Password)
	if err != nil {
		return 0, nil
	}

	lstID, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int64(lstID), nil
}
