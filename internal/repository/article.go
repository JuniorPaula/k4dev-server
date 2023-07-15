package repository

import (
	"database/sql"
	"knowledge-api/internal/models"
)

type article struct {
	DB *sql.DB
}

func NewArticleRepository(db *sql.DB) *article {
	return &article{
		DB: db,
	}
}

func (a *article) CreateArticle(article models.Article) (int64, error) {
	statment, err := a.DB.Prepare("insert into articles (name, description, image_url, content, user_id, category_id) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statment.Close()

	result, err := statment.Exec(article.Name, article.Description, article.ImageURL, article.Content, article.UserID, article.CategoryID)
	if err != nil {
		return 0, err
	}

	lstID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int64(lstID), nil
}
