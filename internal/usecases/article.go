package usecases

import (
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
