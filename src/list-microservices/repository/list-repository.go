package repository

import (
	"context"
	"log"
	domain "traileau-list-microservices/domain/repository"
	model "traileau-list-microservices/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ListRepository struct {
	mongoDB *mongo.Database
}

// DeleteData implements repository.ListRepositoryInterface.
func (l *ListRepository) DeleteData(ctx context.Context, req *string) error {
	panic("unimplemented")
}

// GetAllData implements repository.ListRepositoryInterface.
func (l *ListRepository) GetAllData(ctx context.Context) (lists []model.List, err error) {
	query, err := l.mongoDB.Collection("lists").Find(ctx, bson.D{})
	if err != nil {
		log.Println("error", err)
		return []model.List{}, err
	}
	defer query.Close(ctx)

	listDataList := make([]model.List, 0)
	for query.Next(ctx) {
		var row model.List
		err := query.Decode(&row)
		if err != nil {
			log.Println("error")
		}
		listDataList = append(listDataList, row)
	}

	return listDataList, err
}

// GetData implements repository.ListRepositoryInterface.
func (l *ListRepository) GetData(ctx context.Context, username *string) (list *model.List, err error) {
	panic("unimplemented")
}

// InsertData implements repository.ListRepositoryInterface.
func (l *ListRepository) InsertData(ctx context.Context, req *model.List) error {
	_, err := l.mongoDB.Collection("lists").InsertOne(ctx, req)
	if err != nil {
		return err
	}

	return err
}

// UpdateData implements repository.ListRepositoryInterface.
func (l *ListRepository) UpdateData(ctx context.Context, req *model.List) error {
	panic("unimplemented")
}

func NewProjectRepository(mongo *mongo.Database) domain.ListRepositoryInterface {
	return &ListRepository{
		mongoDB: mongo,
	}
}
