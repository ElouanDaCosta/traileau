package repository

import (
	"context"
	users_model "traileau/src/users/models"
)

type UserRepositoryInterface interface {
	GetAllData(ctx context.Context) (user []users_model.User, err error)
	InsertData(ctx context.Context, req *users_model.User) error
	UpdateData(ctx context.Context, req *users_model.User) error
	DeleteData(ctx context.Context, req *string) error
	GetData(ctx context.Context, username *string) (user *users_model.User, err error)
}
