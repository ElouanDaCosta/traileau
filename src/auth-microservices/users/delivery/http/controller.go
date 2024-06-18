package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	responses "traileau/users/delivery/response"
	usecase "traileau/users/domain/usecase"
	models "traileau/users/models"

	"github.com/go-playground/validator/v10"

	"traileau/users/utils"

	"net/mail"

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

func (uc *UserController) SignUp(ctx *gin.Context) {
	// Initialize the validator
	validate := validator.New(validator.WithRequiredStructEnabled())

	var user models.User

	// Decode the request body to access the data like a json
	decoder := json.NewDecoder(ctx.Request.Body)
	error := decoder.Decode(&user)
	if error != nil {
		fmt.Printf("error %s", error)
		ctx.JSON(501, gin.H{"error": error})
		return
	}
	// Use bcrypt to encrypt the password
	cryptedPassword, errorCryptingPassword := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if errorCryptingPassword != nil {
		fmt.Printf("error %s", error)
		ctx.JSON(501, gin.H{"error": error})
		return
	}

	// Validate the user struct
	errorValidateUser := validate.Struct(&user)

	// Return 400 if missing user input
	if errorValidateUser != nil {
		fmt.Println("validation failed 2")
		ctx.JSON(400, gin.H{"error": error})
		return
	}

	validEmail, mailError := mail.ParseAddress(user.Email)

	if mailError != nil {
		fmt.Println("Email not valid")
		ctx.JSON(400, gin.H{"error": error})
		return
	}

	_, userExistingError := uc.UserUseCase.GetUser(ctx, &user.Email)

	if userExistingError == nil {
		fmt.Println("User already exist")
		ctx.JSON(400, gin.H{"error": error})
		return
	}

	newUser := models.User{
		Username: user.Username,
		Email:    validEmail.Address,
		Password: string(cryptedPassword),
	}

	// Insert the new user in db
	err := uc.UserUseCase.CreateUser(ctx, &newUser)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success"})
}

func (uc *UserController) SignIn(ctx *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read the body.",
		})
	}

	user, userError := uc.UserUseCase.GetUser(ctx, &body.Email)

	if userError != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
	}

	checkUserPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	token, tokenError := utils.CreateToken(body.Email)

	if tokenError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Generating token",
		})
	}

	if checkUserPassword != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong password",
		})
	} else {
		ctx.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: token})
	}

}

func (uc *UserController) GetAll(ctx *gin.Context) {
	users, err := uc.UserUseCase.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: users})
}

func (uc *UserController) GetOne(ctx *gin.Context) {
	email, err := ctx.GetQuery("email")

	if !err {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read the query",
		})
	}

	user, userError := uc.UserUseCase.GetUser(ctx, &email)

	if userError != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
	} else {
		ctx.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: user})
	}
}
