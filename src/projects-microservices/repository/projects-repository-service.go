package repository

import (
	"context"
	domain "traileau-projects-microservices/domain/repository"
	model "traileau-projects-microservices/models"

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
	panic("unimplemented")
}

// GetData implements repository.ProjectRepositoryInterface.
func (p *ProjectRepository) GetData(ctx context.Context, username *string) (user *model.Project, err error) {
	panic("unimplemented")
}

// InsertData implements repository.ProjectRepositoryInterface.
func (p *ProjectRepository) InsertData(ctx context.Context, req *model.Project) error {
	panic("unimplemented")
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
