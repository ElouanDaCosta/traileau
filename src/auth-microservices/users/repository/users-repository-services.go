package repository

import (
	"context"
	"log"
	domain "traileau/users/domain/repository"
	models "traileau/users/models"

	"go.mongodb.org/mongo-driver/bson"
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
func (c UserRepository) GetAllData(ctx context.Context) (userResp []models.User, err error) {
	query, err := c.mongoDB.Collection("users").Find(ctx, bson.D{})
	if err != nil {
		log.Println("error", err)
		return []models.User{}, err
	}
	defer query.Close(ctx)

	listDataUser := make([]models.User, 0)
	for query.Next(ctx) {
		var row models.User
		err := query.Decode(&row)
		if err != nil {
			log.Println("error")
		}
		listDataUser = append(listDataUser, row)
	}

	return listDataUser, err
}

// GetData implements repository.UserRepositoryInterface.
func (c *UserRepository) GetData(ctx context.Context, username *string) (user *models.User, err error) {
	var result struct {
		Name     string `bson:"username"`
		Email    string `bson:"email"`
		Password string `bson:"password"`
	}
	collection := c.mongoDB.Collection("users")
	filter := bson.D{{Key: "email", Value: username}}

	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	user = &models.User{
		Username: result.Name,
		Email:    result.Email,
		Password: result.Password,
	}

	return user, err
}

// UpdateData implements repository.UserRepositoryInterface.
func (c *UserRepository) UpdateData(ctx context.Context, req *models.User) error {
	panic("unimplemented")
}

func (c UserRepository) InsertData(ctx context.Context, req *models.User) error {

	_, err := c.mongoDB.Collection("users").InsertOne(ctx, req)
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
