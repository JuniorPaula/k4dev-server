package usecases

import (
	"knowledge-api/internal/database"
	"knowledge-api/internal/models"
	"knowledge-api/internal/repository"
)

func CreateCategoryUSecase(c models.Category) (models.Category, error) {
	if err := c.HanlderCategory(); err != nil {
		return models.Category{}, err
	}

	db, err := database.Connect_MySQL()
	if err != nil {
		return models.Category{}, err
	}
	defer db.Close()

	repo := repository.NewCategoryRepository(db)
	c.ID, err = repo.CreateCategory(c)
	if err != nil {
		return models.Category{}, err
	}

	return c, nil
}
