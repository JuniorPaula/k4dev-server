package repository

import (
	"context"
	"errors"
	"knowledge-api/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (s *stat) InsertStat(stat models.Stat) (primitive.ObjectID, error) {
	result, err := s.Mongo.Collection("stats").InsertOne(context.Background(), stat)
	if err != nil {
		return primitive.NilObjectID, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid, nil
	}

	return primitive.NilObjectID, nil
}

func (s *stat) UpdateStat(stat models.Stat) error {
	_, err := s.Mongo.Collection("stats").UpdateOne(context.Background(), bson.M{}, bson.M{"$set": stat})
	if err != nil {
		return err
	}

	return nil
}
