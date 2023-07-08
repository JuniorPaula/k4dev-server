package models

import (
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

type User struct {
	ID       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Admin    bool   `json:"admin,omitempty"`
}

func (u *User) HanlderUser(step string) error {
	if err := u.validator(step); err != nil {
		return err
	}
	if err := u.formatFields(step); err != nil {
		return err
	}
	return nil
}

func (u *User) validator(step string) error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("invalid email")
	}

	if step == "create" && u.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func (u *User) formatFields(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	if step == "create" {
		u.Password = strings.TrimSpace(u.Password)
	}
	return nil
}
