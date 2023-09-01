package usecases

import (
	"knowledge-api/internal/database"
	"knowledge-api/internal/models"
	"knowledge-api/internal/repository"
	"log"
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

func StatScheduleUsecase() {
	mongoDB, err := database.Connect_MongoDB()
	if err != nil {
		log.Printf("[ERROR] connect to mongoDB: %v", err)
	}

	db, err := database.Connect_MySQL()
	if err != nil {
		log.Printf("[ERROR] connect to mysql: %v", err)
	}
	defer db.Close()

	statRepo := repository.NewStatRepository(mongoDB)
	usersRepo := repository.NewUsersRepository(db)
	articleRepo := repository.NewArticleRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)

	users, err := usersRepo.FindAllUsers()
	if err != nil {
		log.Printf("[ERROR] to find users: %v", err)
		return
	}

	articlesCount, err := articleRepo.CountArticles()
	if err != nil {
		log.Printf("[ERROR] to count articles: %v", err)
		return
	}

	categories, err := categoryRepo.FindAllCategories()
	if err != nil {
		log.Printf("[ERROR] to find categories: %v", err)
		return
	}

	ltsStats, _ := statRepo.FindStat()

	usersCount := len(users)
	categoriesCount := len(categories)

	if (ltsStats == models.Stat{}) {
		_, err = statRepo.InsertStat(models.Stat{
			Article:   articlesCount,
			Category:  categoriesCount,
			Users:     usersCount,
			CreatedAt: time.Now(),
		})
		if err != nil {
			log.Printf("[ERROR] to insert stats: %v", err)
			return
		}

		log.Println("[SUCCESS] stats inserted")
	}

	if ltsStats.Article != articlesCount || ltsStats.Category != categoriesCount || ltsStats.Users != usersCount {
		statRepo.UpdateStat(models.Stat{
			Article:   articlesCount,
			Category:  categoriesCount,
			Users:     usersCount,
			CreatedAt: time.Now(),
		})

		log.Println("[SUCCESS] stats updated")

	} else {
		log.Println("[SUCCESS] stats not updated")
	}

}
