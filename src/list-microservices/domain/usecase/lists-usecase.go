package usecase

import (
	"context"
	model "traileau-list-microservices/models"
)

type ListUseCase interface {
	CreateList(ctx context.Context, req *model.List) error
	GetList(ctx context.Context, req *string) (*model.List, error)
	GetAll(ctx context.Context) ([]model.List, error)
	UpdateList(ctx context.Context, req *model.List) error
	DeleteList(ctx context.Context, req *string) error
}
