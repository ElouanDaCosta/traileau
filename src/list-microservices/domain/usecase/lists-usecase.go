package usecase

import (
	"context"
	model "traileau-list-microservices/models"
)

type ListUseCase interface {
	CreateProject(ctx context.Context, req *model.List) error
	GetProject(ctx context.Context, req *string) (*model.List, error)
	GetAll(ctx context.Context) ([]model.List, error)
	UpdateProject(ctx context.Context, req *model.List) error
	DeleteProject(ctx context.Context, req *string) error
}
