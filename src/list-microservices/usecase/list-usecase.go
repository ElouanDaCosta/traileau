package usecase

import (
	"context"
	domain_repository "traileau-list-microservices/domain/repository"
	domain_usecase "traileau-list-microservices/domain/usecase"
	model "traileau-list-microservices/models"
)

type ListUseCase struct {
	listRepo domain_repository.ListRepositoryInterface
	ctx      context.Context
}

// CreateProject implements usecase.ListUseCase.
func (l *ListUseCase) CreateProject(ctx context.Context, req *model.List) error {
	panic("unimplemented")
}

// DeleteProject implements usecase.ListUseCase.
func (l *ListUseCase) DeleteProject(ctx context.Context, req *string) error {
	panic("unimplemented")
}

// GetAll implements usecase.ListUseCase.
func (l *ListUseCase) GetAll(ctx context.Context) ([]model.List, error) {
	panic("unimplemented")
}

// GetProject implements usecase.ListUseCase.
func (l *ListUseCase) GetProject(ctx context.Context, req *string) (*model.List, error) {
	panic("unimplemented")
}

// UpdateProject implements usecase.ListUseCase.
func (l *ListUseCase) UpdateProject(ctx context.Context, req *model.List) error {
	panic("unimplemented")
}

func NewProjectUsecase(listRepo domain_repository.ListRepositoryInterface, ctx context.Context) domain_usecase.ListUseCase {
	return &ListUseCase{
		listRepo: listRepo,
		ctx:      ctx,
	}
}
