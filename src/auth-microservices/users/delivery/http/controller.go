package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	responses "traileau/users/delivery/response"
	usecase "traileau/users/domain/usecase"
	models "traileau/users/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	UserUseCase usecase.UserUsecase
}

func New(userservice usecase.UserUsecase) UserController {
	return UserController{
		UserUseCase: userservice,
	}
}

func (uc *UserController) Register(ctx *gin.Context) {
	var user models.User
	decoder := json.NewDecoder(ctx.Request.Body)
	error := decoder.Decode(&user)
	if error != nil {
		fmt.Printf("error %s", error)
		ctx.JSON(501, gin.H{"error": error})
	}
	//fmt.Printf("Decode Body %v\n\n", user)
	if user.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Email is required"})
		return
	}
	cryptedPassword, errorCryptingPassword := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if errorCryptingPassword != nil {
		fmt.Printf("error %s", error)
		ctx.JSON(501, gin.H{"error": error})
	}

	newUser := models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: string(cryptedPassword),
	}

	err := uc.UserUseCase.CreateUser(ctx, &newUser)
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
