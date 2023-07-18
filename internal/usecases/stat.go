package usecases

import (
	"fmt"
	"knowledge-api/internal/database"
	"knowledge-api/internal/models"
	"knowledge-api/internal/repository"
	"time"
)

func FindStatUsecase() (models.Stat, error) {
	mongoDB, err := database.Connect_MongoDB()
	if err != nil {
		return models.Stat{}, err
	}

	statRepo := repository.NewStatRepository(mongoDB)

	stats, err := statRepo.FindStat()
	if err != nil {
		return models.Stat{}, err
	}

	return stats, nil
}

func StatScheduleUsecase() error {
	mongoDB, err := database.Connect_MongoDB()
	if err != nil {
		return err
	}

	db, err := database.Connect_MySQL()
	if err != nil {
		return err
	}

	statRepo := repository.NewStatRepository(mongoDB)
	usersRepo := repository.NewUsersRepository(db)
	articleRepo := repository.NewArticleRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)

	users, err := usersRepo.FindAllUsers()
	if err != nil {
		return err
	}

	articles, err := articleRepo.FindAllArticles(10, 10)
	if err != nil {
		return err
	}

	categories, err := categoryRepo.FindAllCategories()
	if err != nil {
		return err
	}

	ltsStats, _ := statRepo.FindStat()

	usersCount := len(users)
	articlesCount := len(articles)
	categoriesCount := len(categories)

	_, err = statRepo.InsertStat(models.Stat{
		Article:   articlesCount,
		Category:  categoriesCount,
		Users:     usersCount,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	if ltsStats.Article != articlesCount || ltsStats.Category != categoriesCount || ltsStats.Users != usersCount {
		statRepo.UpdateStat(models.Stat{
			Article:   articlesCount,
			Category:  categoriesCount,
			Users:     usersCount,
			CreatedAt: time.Now(),
		})

		fmt.Println("stats updated on success")
	}

	fmt.Println("nothing to update")
	return nil
}
