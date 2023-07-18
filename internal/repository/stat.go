package repository

import (
	"context"
	"errors"
	"knowledge-api/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type stat struct {
	Mongo *mongo.Database
}

func NewStatRepository(mongo *mongo.Database) *stat {
	return &stat{
		Mongo: mongo,
	}
}

func (s *stat) FindStat() (models.Stat, error) {
	var stats models.Stat
	err := s.Mongo.Collection("stats").FindOne(context.Background(), bson.D{}).Decode(&stats)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Stat{
				Article:  0,
				Category: 0,
				Users:    0,
			}, errors.New("stat not data")
		} else {
			return models.Stat{}, err
		}
	}

	return stats, nil
}
