package http

import (
	"net/http"
	responses "traileau/users/delivery/response"
	usecase "traileau/users/domain/usecase"
	models "traileau/users/models"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUseCase usecase.UserUsecase
}

func New(userservice usecase.UserUsecase) UserController {
	return UserController{
		UserUseCase: userservice,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserUseCase.CreateUser(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success"})
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	users, err := uc.UserUseCase.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: users})
}
