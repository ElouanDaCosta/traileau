package repository

import (
	project_model "boards-projects-microservices/models"
	"context"
)

type ProjectRepositoryInterface interface {
	GetAllData(ctx context.Context) (user []project_model.Project, err error)
	InsertData(ctx context.Context, req *project_model.Project) error
	UpdateData(ctx context.Context, req *project_model.Project) error
	DeleteData(ctx context.Context, req *string) error
	GetData(ctx context.Context, username *string) (user *project_model.Project, err error)
}
