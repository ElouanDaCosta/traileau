package usecase

import (
	"context"
	"log"
	domain_repository "traileau-list-microservices/domain/repository"
	domain_usecase "traileau-list-microservices/domain/usecase"
	model "traileau-list-microservices/models"
)

type ListUseCase struct {
	listRepo domain_repository.ListRepositoryInterface
	ctx      context.Context
}

// GetAll implements usecase.ListUseCase.
func (l *ListUseCase) GetAll(ctx context.Context) ([]model.List, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	list, err := l.listRepo.GetAllData(ctx)
	if err != nil {
		log.Println("failed to show data project with default log")
		return list, err
	}

	return list, err
}

// CreateList implements usecase.ListUseCase.
func (l *ListUseCase) CreateList(ctx context.Context, req *model.List) error {
	if ctx == nil {
		ctx = context.Background()
	}

	err := l.listRepo.InsertData(ctx, req)
	if err != nil {
		return err
	}
	log.Println("Successfully Inserted List Data")

	return err
}

// DeleteList implements usecase.ListUseCase.
func (l *ListUseCase) DeleteList(ctx context.Context, req *string) error {
	if ctx == nil {
		ctx = context.Background()
	}

	err := l.listRepo.DeleteData(ctx, req)
	if err != nil {
		return err
	}
	log.Println("Successfully Deleted List Data")

	return err
}

// GetList implements usecase.ListUseCase.
func (l *ListUseCase) GetList(ctx context.Context, name *string) (*model.List, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	list, err := l.listRepo.GetData(ctx, name)
	if err != nil {
		log.Println("failed to show list data with default log")
		return list, err
	}

	return list, err
}

// UpdateList implements usecase.ListUseCase.
func (l *ListUseCase) UpdateList(ctx context.Context, req *model.List) error {
	panic("unimplemented")
}

func NewProjectUsecase(listRepo domain_repository.ListRepositoryInterface, ctx context.Context) domain_usecase.ListUseCase {
	return &ListUseCase{
		listRepo: listRepo,
		ctx:      ctx,
	}
}
