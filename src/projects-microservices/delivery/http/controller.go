package http

import (
	usecase "traileau-projects-microservices/domain/usecase"
)

type ProjectController struct {
	ProjectUseCase usecase.ProjectUseCase
}

func New(projectservice usecase.ProjectUseCase) ProjectController {
	return ProjectController{
		ProjectUseCase: projectservice,
	}
}
