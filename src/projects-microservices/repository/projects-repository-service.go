package repository

import (
	"context"
	"log"
	domain "traileau-projects-microservices/domain/repository"
	model "traileau-projects-microservices/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectRepository struct {
	mongoDB *mongo.Database
}

// DeleteData implements repository.ProjectRepositoryInterface.
func (p *ProjectRepository) DeleteData(ctx context.Context, req *string) error {
	panic("unimplemented")
}

// GetAllData implements repository.ProjectRepositoryInterface.
func (p *ProjectRepository) GetAllData(ctx context.Context) (user []model.Project, err error) {
	query, err := p.mongoDB.Collection("projects").Find(ctx, bson.D{})
	if err != nil {
		log.Println("error", err)
		return []model.Project{}, err
	}
	defer query.Close(ctx)

	listDataProject := make([]model.Project, 0)
	for query.Next(ctx) {
		var row model.Project
		err := query.Decode(&row)
		if err != nil {
			log.Println("error")
		}
		listDataProject = append(listDataProject, row)
	}

	return listDataProject, err
}

// GetData implements repository.ProjectRepositoryInterface.
func (p *ProjectRepository) GetData(ctx context.Context, username *string) (user *model.Project, err error) {
	panic("unimplemented")
}

// InsertData implements repository.ProjectRepositoryInterface.
func (p *ProjectRepository) InsertData(ctx context.Context, req *model.Project) error {
	_, err := p.mongoDB.Collection("projects").InsertOne(ctx, req)
	if err != nil {
		log.Println("error")
	}

	return err
}

// UpdateData implements repository.ProjectRepositoryInterface.
func (p *ProjectRepository) UpdateData(ctx context.Context, req *model.Project) error {
	panic("unimplemented")
}

func NewProjectRepository(mongo *mongo.Database) domain.ProjectRepositoryInterface {
	return &ProjectRepository{
		mongoDB: mongo,
	}
}
