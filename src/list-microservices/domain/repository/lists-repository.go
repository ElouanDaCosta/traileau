package repository

import (
	"context"
	list_model "traileau-list-microservices/models"
)

type ListRepositoryInterface interface {
	GetAllData(ctx context.Context) (lists []list_model.List, err error)
	InsertData(ctx context.Context, req *list_model.List) error
	UpdateData(ctx context.Context, req *list_model.List) error
	DeleteData(ctx context.Context, req *string) error
	GetData(ctx context.Context, username *string) (list *list_model.List, err error)
}
