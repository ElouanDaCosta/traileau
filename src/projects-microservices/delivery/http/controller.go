package http

import (
	"net/http"
	responses "traileau-projects-microservices/delivery/response"
	usecase "traileau-projects-microservices/domain/usecase"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	ProjectUseCase usecase.ProjectUseCase
}

func New(projectservice usecase.ProjectUseCase) ProjectController {
	return ProjectController{
		ProjectUseCase: projectservice,
	}
}

func (pc *ProjectController) GetAll(ctx *gin.Context) {
	projects, err := pc.ProjectUseCase.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, responses.ProjectResponse{Status: http.StatusOK, Message: "success", Data: projects})
}
