package models

import (
	"errors"
	"fmt"
)

type Article struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url,omitempty"`
	Content     string `json:"content"`
	UserID      int64  `json:"user_id"`
	CategoryID  int64  `json:"category_id"`
}

func (a *Article) HandeArticles() error {
	requiredFields := map[string]string{
		"name":        a.Name,
		"description": a.Description,
		"content":     string(a.Content),
		"user_id":     fmt.Sprint(a.UserID),
		"category_id": fmt.Sprint(a.CategoryID),
	}

	maxLengthFields := map[string]int{
		"description": 1000,
	}

	for field, value := range requiredFields {
		if value == "" {
			return errors.New("the " + field + " field is required")
		}
	}

	for field, maxLength := range maxLengthFields {
		if len(requiredFields[field]) > maxLength {
			return errors.New("the " + field + " field must be less than " + fmt.Sprint(maxLength) + " characters")
		}
	}

	return nil
}
