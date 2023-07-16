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

func (a *article) UpdateArticle(id int64, article models.Article) error {
	statment, err := a.DB.Prepare("update articles set name = ?, description = ?, image_url = ?, content = ?, category_id = ? where id = ?")
	if err != nil {
		return err
	}
	defer statment.Close()

	if _, err := statment.Exec(article.Name, article.Description, article.ImageURL, article.Content, article.CategoryID, id); err != nil {
		return err
	}

	return nil
}

func (a *article) FindArticleByID(id int64) (models.Article, error) {
	var article models.Article

	statment, err := a.DB.Prepare("select id, name, description, image_url, content, user_id, category_id from articles where id = ?")
	if err != nil {
		return article, err
	}
	defer statment.Close()

	if err := statment.QueryRow(id).Scan(&article.ID, &article.Name, &article.Description, &article.ImageURL, &article.Content, &article.UserID, &article.CategoryID); err != nil {
		return article, err
	}

	return article, nil
}
