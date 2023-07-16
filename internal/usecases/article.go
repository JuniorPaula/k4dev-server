package usecases

import (
	"errors"
	"knowledge-api/internal/database"
	"knowledge-api/internal/models"
	"knowledge-api/internal/repository"
)

func CreateArticleUsecase(article models.Article) (models.Article, error) {
	if err := article.HandeArticles(); err != nil {
		return article, err
	}

	db, err := database.Connect_MySQL()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	repositories := repository.NewArticleRepository(db)

	articleID, err := repositories.CreateArticle(article)
	if err != nil {
		return article, err
	}

	article.ID = articleID

	return article, nil
}

func FindAllArticlesUsecase(page, pageSize int) ([]models.Article, error) {
	db, err := database.Connect_MySQL()
	if err != nil {
		return []models.Article{}, err
	}
	defer db.Close()

	articleRepo := repository.NewArticleRepository(db)
	if err != nil {
		return []models.Article{}, err
	}

	articles, err := articleRepo.FindAllArticles(page, pageSize)
	if err != nil {
		return []models.Article{}, err
	}

	return articles, nil
}

func FindArticleByIDUsecase(id int64) (models.Article, error) {
	db, err := database.Connect_MySQL()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	articleRepo := repository.NewArticleRepository(db)

	article, err := articleRepo.FindArticleByID(id)
	if err != nil {
		return article, err
	}

	article.Content = string(article.Content)

	return article, nil
}

func UpdateArticlesUsecase(id, userIDInToken int64, a models.Article) error {
	if err := a.HandeArticles(); err != nil {
		return err
	}

	db, err := database.Connect_MySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	articleRepo := repository.NewArticleRepository(db)
	articleFromDB, err := articleRepo.FindArticleByID(id)
	if err != nil {
		return err
	}

	if userIDInToken != articleFromDB.UserID {
		return errors.New("you are not the owner of this article")
	}

	if err := articleRepo.UpdateArticle(id, a); err != nil {
		return err
	}

	return nil
}

func DeleteArticleUsecase(id, userIDInToken int64) error {
	db, err := database.Connect_MySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	articleRepo := repository.NewArticleRepository(db)
	articleFromDB, err := articleRepo.FindArticleByID(id)
	if err != nil {
		return err
	}

	userRepo := repository.NewUsersRepository(db)
	userFromDB, err := userRepo.FindUserByID(userIDInToken)

	if err != nil {
		return err
	}

	if !userFromDB.Admin && userIDInToken != articleFromDB.UserID {
		return errors.New("you are not the owner of this article")
	}

	if err := articleRepo.DeleteArticle(id); err != nil {
		return err
	}

	return nil
}

func CountArticlesUsecase() (int, error) {
	db, err := database.Connect_MySQL()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	articleRepo := repository.NewArticleRepository(db)

	count, err := articleRepo.CountArticles()
	if err != nil {
		return 0, err
	}

	return count, nil
}
