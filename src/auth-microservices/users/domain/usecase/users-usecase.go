package usecase

import (
	"context"
	model "traileau/users/models"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, req *model.User) error
	GetUser(ctx context.Context, req *string) (*model.User, error)
	GetAll(ctx context.Context) ([]model.User, error)
	UpdateUser(ctx context.Context, req *model.User) error
	DeleteUser(ctx context.Context, req *string) error
}
