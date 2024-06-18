package usecase

import (
	"context"
	model "traileau-projects-microservices/models"
)

type ProjectUseCase interface {
	CreateProject(ctx context.Context, req *model.Project) error
	GetProject(ctx context.Context, req *string) (*model.Project, error)
	GetAll(ctx context.Context) ([]model.Project, error)
	UpdateProject(ctx context.Context, req *model.Project) error
	DeleteProject(ctx context.Context, req *string) error
}
