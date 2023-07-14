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

func (c *category) UpdateCategory(id int64, category models.Category) error {
	statment, err := c.DB.Prepare("update categories set name = ? where id = ?")
	if err != nil {
		return err
	}
	defer statment.Close()

	if _, err := statment.Exec(category.Name, id); err != nil {
		return nil
	}

	return nil
}

func (c *category) DeleteCategory(id int64) error {
	statment, err := c.DB.Prepare("delete from categories where id = ?")
	if err != nil {
		return err
	}
	defer statment.Close()

	if _, err := statment.Exec(id); err != nil {
		return err
	}

	return nil
}

func (c *category) VerifyCategoryHasParentID(id int64) (bool, error) {
	statment, err := c.DB.Prepare("select * from categories where parent_id = ?")
	if err != nil {
		return false, err
	}
	defer statment.Close()

	rows, err := statment.Query(id)
	if err != nil {
		return false, err
	}

	if rows.Next() {
		return true, nil
	}

	return false, nil
}
