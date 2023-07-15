package usecases

import (
	"errors"
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

func FindAllCategoriesUsecase() ([]models.Category, error) {
	db, err := database.Connect_MySQL()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	repo := repository.NewCategoryRepository(db)
	categories, err := repo.FindAllCategories()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func UpdateCategoryUsecase(id, userIDInToken int64, c models.Category) error {
	if err := c.HanlderCategory(); err != nil {
		return err
	}

	db, err := database.Connect_MySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	repo := repository.NewCategoryRepository(db)
	userRepo := repository.NewUsersRepository(db)

	userFromDB, err := userRepo.FindUserByID(userIDInToken)
	if err != nil {
		return err
	}

	if !userFromDB.Admin {
		return errors.New("user not have permission")
	}

	err = repo.UpdateCategory(id, c)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCategoryUsecase(id, userIDInToken int64, c models.Category) error {
	db, err := database.Connect_MySQL()
	if err != nil {
		return errors.New("error connect database")
	}
	defer db.Close()

	categoryRepo := repository.NewCategoryRepository(db)
	userRepo := repository.NewUsersRepository(db)

	userFromDB, err := userRepo.FindUserByID(userIDInToken)
	if err != nil {
		return err
	}

	if !userFromDB.Admin {
		return errors.New("user not have permission")
	}

	ok, err := categoryRepo.VerifyCategoryHasParentID(id)
	if err != nil {
		return err
	}
	if ok {
		return errors.New("category has sub category, please delete it first")
	}

	// TODO - implements the same validation for articles

	err = categoryRepo.DeleteCategory(id)
	if err != nil {
		return err
	}

	return nil

}
