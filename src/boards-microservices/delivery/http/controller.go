package http

import (
	responses "boards-projects-microservices/delivery/response"
	usecase "boards-projects-microservices/domain/usecase"
	helper "boards-projects-microservices/helpers"
	model "boards-projects-microservices/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func (pc *ProjectController) CreateProject(ctx *gin.Context) {
	// Initialize the validator
	validate := validator.New(validator.WithRequiredStructEnabled())

	userToken, userErr := helper.GetTokenData(ctx)

	if userErr != nil {
		fmt.Printf("error %s", userErr)
		ctx.JSON(501, gin.H{"error": userErr})
		return
	}

	existingUser, existingUserError := GetUser(userToken)

	if existingUserError != nil {
		fmt.Printf("error %s", existingUserError)
		ctx.JSON(501, gin.H{"error": existingUserError})
		return
	}

	if existingUser.StatusCode == 404 {
		fmt.Printf("error %s", existingUserError)
		ctx.JSON(401, gin.H{"error": "User not found from the current session"})
		ctx.Abort()
		return
	}

	var project model.Project

	// Decode the request body to access the data like a json
	decoder := json.NewDecoder(ctx.Request.Body)
	error := decoder.Decode(&project)
	if error != nil {
		fmt.Printf("error %s", error)
		ctx.JSON(501, gin.H{"error": error})
		return
	}

	// Validate the user struct
	errorValidateProject := validate.Struct(&project)

	// Return 400 if missing user input
	if errorValidateProject != nil {
		fmt.Println("validation failed for project struct")
		fmt.Println(errorValidateProject)
		ctx.JSON(400, gin.H{"error": "Failed to validate the project structure"})
		return
	}

	newProject := model.Project{
		Name:        project.Name,
		Description: project.Description,
		Author:      userToken,
	}

	err := pc.ProjectUseCase.CreateProject(ctx, &newProject)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, responses.ProjectResponse{Status: http.StatusOK, Message: "success", Data: newProject})
}
