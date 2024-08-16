package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	response "traileau-list-microservices/delivery/response"
	usecase "traileau-list-microservices/domain/usecase"
	model "traileau-list-microservices/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ListController struct {
	ListUseCase usecase.ListUseCase
}

func New(listservice usecase.ListUseCase) ListController {
	return ListController{
		ListUseCase: listservice,
	}
}

func (lc *ListController) GetAll(ctx *gin.Context) {
	lists, err := lc.ListUseCase.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response.ListResponse{Status: http.StatusOK, Message: "success", Data: lists})
}

func (lc *ListController) CreateList(ctx *gin.Context) {
	// Initialize the validator
	validate := validator.New(validator.WithRequiredStructEnabled())

	var list model.List

	// Decode the request body to access the data like a json
	decoder := json.NewDecoder(ctx.Request.Body)
	error := decoder.Decode(&list)
	if error != nil {
		fmt.Printf("error %s", error)
		ctx.JSON(501, gin.H{"error": error})
		return
	}

	errorValidateList := validate.Struct(&list)

	if errorValidateList != nil {
		ctx.JSON(400, gin.H{"error": "Failed to validate the list structure"})
		return
	}

	newList := model.List{
		Id:       primitive.NewObjectID(),
		Name:     list.Name,
		Position: list.Position,
	}

	err := lc.ListUseCase.CreateList(ctx, &newList)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.ListResponse{Status: http.StatusOK, Message: "success", Data: newList})
}
