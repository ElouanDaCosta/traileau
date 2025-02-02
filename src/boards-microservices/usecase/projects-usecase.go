package usecase

import (
	domain1 "boards-projects-microservices/domain/repository"
	domain2 "boards-projects-microservices/domain/usecase"
	projects_model "boards-projects-microservices/models"
	"context"
	"log"
)

type ProjectServiceImpl struct {
	projectRepo domain1.ProjectRepositoryInterface
	ctx         context.Context
}

// CreateProject implements usecase.ProjectUseCase.
func (p *ProjectServiceImpl) CreateProject(ctx context.Context, req *projects_model.Project) error {
	if ctx == nil {
		ctx = context.Background()
	}

	err := p.projectRepo.InsertData(ctx, req)
	if err != nil {
		return err
	}
	log.Println("Successfully Inserted Project Data")

	return nil
}

// DeleteProject implements usecase.ProjectUseCase.
func (p *ProjectServiceImpl) DeleteProject(ctx context.Context, req *string) error {
	panic("unimplemented")
}

// GetAll implements usecase.ProjectUseCase.
func (p *ProjectServiceImpl) GetAll(ctx context.Context) ([]projects_model.Project, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	list, err := p.projectRepo.GetAllData(ctx)
	if err != nil {
		log.Println("failed to show data project with default log")
		return list, err
	}

	return list, err
}

// GetProject implements usecase.ProjectUseCase.
func (p *ProjectServiceImpl) GetProject(ctx context.Context, req *string) (*projects_model.Project, error) {
	panic("unimplemented")
}

// UpdateProject implements usecase.ProjectUseCase.
func (p *ProjectServiceImpl) UpdateProject(ctx context.Context, req *projects_model.Project) error {
	panic("unimplemented")
}

func NewProjectUsecase(projectRepo domain1.ProjectRepositoryInterface, ctx context.Context) domain2.ProjectUseCase {
	return &ProjectServiceImpl{
		projectRepo: projectRepo,
		ctx:         ctx,
	}
}
