package repository

import (
	"database/sql"
	"knowledge-api/internal/models"
	"strings"
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
	query := "insert into categories (name"
	values := []interface{}{category.Name}

	if category.ParentID != 0 {
		query += ", parent_id"
		values = append(values, category.ParentID)
	}

	query += ") values (?" + strings.Repeat(", ?", len(values)-1) + ")"

	statment, err := c.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer statment.Close()

	result, err := statment.Exec(values...)
	if err != nil {
		return 0, err
	}

	lstID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int64(lstID), nil

}
