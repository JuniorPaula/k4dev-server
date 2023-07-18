package usecases

import (
	"knowledge-api/internal/database"
	"knowledge-api/internal/models"
	"knowledge-api/internal/repository"
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
