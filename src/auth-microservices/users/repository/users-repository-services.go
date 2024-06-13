package repository

import (
	"context"
	"log"
	domain "traileau/users/domain/repository"
	models "traileau/users/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	mongoDB *mongo.Database
}

// DeleteData implements repository.UserRepositoryInterface.
func (c *UserRepository) DeleteData(ctx context.Context, req *string) error {
	panic("unimplemented")
}

// GetAllData implements repository.UserRepositoryInterface.
func (c *UserRepository) GetAllData(ctx context.Context) (user []models.User, err error) {
	panic("unimplemented")
}

// GetData implements repository.UserRepositoryInterface.
func (c *UserRepository) GetData(ctx context.Context, username *string) (user *models.User, err error) {
	panic("unimplemented")
}

// UpdateData implements repository.UserRepositoryInterface.
func (c *UserRepository) UpdateData(ctx context.Context, req *models.User) error {
	panic("unimplemented")
}

func (c UserRepository) InsertData(ctx context.Context, req *models.User) error {

	_, err := c.mongoDB.Collection("students").InsertOne(ctx, req)
	if err != nil {
		log.Println("error")
	}

	return err
}

func NewUserRepository(mongo *mongo.Database) domain.UserRepositoryInterface {
	return &UserRepository{
		mongoDB: mongo,
	}
}
