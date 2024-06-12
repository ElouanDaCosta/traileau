package repository

import (
	"context"
	"log"
	models "traileau/src/users/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	mongoDB *mongo.Database
}

func (c UserRepository) InsertData(ctx context.Context, req *models.User) error {

	_, err := c.mongoDB.Collection("students").InsertOne(ctx, req)
	if err != nil {
		log.Println("error")
	}

	return err
}
