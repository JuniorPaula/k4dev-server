package repository

import (
	"database/sql"
	"knowledge-api/internal/models"
)

type category struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) *category {
	return &category{
		DB: db,
	}
}

func (c *category) CreateCategory(category models.Category) (int64, error) {
	statment, err := c.DB.Prepare("insert into categories (name, parent_id) values (?, ?)")
	if err != nil {
		return 0, err
	}
	defer statment.Close()

	result, err := statment.Exec(category.Name, category.ParentID)
	if err != nil {
		return 0, err
	}

	lstID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int64(lstID), nil
}
