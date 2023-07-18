package repository

import (
	"database/sql"
	"knowledge-api/internal/models"
	"strings"
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

func (a *article) FindAllArticles(page, pageSize int) ([]models.Article, error) {
	offset := (page - 1) * pageSize

	query := "select id, name, description, image_url, content, user_id, category_id from articles limit ? offset ?"

	rows, err := a.DB.Query(query, pageSize, offset)
	if err != nil {
		return []models.Article{}, err
	}
	defer rows.Close()

	var articles []models.Article

	for rows.Next() {
		var article models.Article

		if err := rows.Scan(&article.ID, &article.Name, &article.Description, &article.ImageURL, &article.Content, &article.UserID, &article.CategoryID); err != nil {
			return []models.Article{}, err
		}

		articles = append(articles, article)
	}

	if err := rows.Err(); err != nil {
		return []models.Article{}, err
	}

	return articles, nil
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

func (a *article) DeleteArticle(id int64) error {
	statment, err := a.DB.Prepare("delete from articles where id = ?")
	if err != nil {
		return err
	}
	defer statment.Close()

	if _, err := statment.Exec(id); err != nil {
		return err
	}

	return nil
}

func (a *article) CountArticles() (int, error) {
	var count int

	query := "select count(id) from articles"

	if err := a.DB.QueryRow(query).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (a *article) FindCategoryWithChildren(ids []int, page, limit int) ([]models.Article, error) {
	offset := (page - 1) * limit

	query := `
	SELECT a.id, a.name, a.description, a.content, a.user_id, a.category_id, a.image_url, u.name AS author
	FROM articles a
	JOIN users u ON u.id = a.user_id
	WHERE a.category_id IN (` + strings.Join(strings.Split(strings.Repeat("?", len(ids)), ""), ", ") + `)
	ORDER BY a.id DESC
	LIMIT ? OFFSET ?
	`

	values := make([]interface{}, 0, len(ids)+2)
	for _, id := range ids {
		values = append(values, id)
	}
	values = append(values, limit, offset)

	rows, err := a.DB.Query(query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var article models.Article

		if err := rows.Scan(&article.ID, &article.Name, &article.Description, &article.Content, &article.UserID, &article.CategoryID, &article.ImageURL, &article.Author); err != nil {
			return nil, err
		}

		article.Content = string(article.Content)
		articles = append(articles, article)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}

func (a *article) VerifyArticleHasCategoryID(categoryID int64) (bool, error) {
	statment, err := a.DB.Prepare("select * from articles where category_id = ?")
	if err != nil {
		return false, err
	}
	defer statment.Close()

	rows, err := statment.Query(categoryID)
	if err != nil {
		return false, err
	}

	if rows.Next() {
		return true, nil
	}

	return false, nil
}
