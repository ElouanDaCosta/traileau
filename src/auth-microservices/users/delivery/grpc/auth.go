package auth

import (
	"context"
	"fmt"
	usecase "traileau-auth-microservices/users/domain/usecase"
	model "traileau-auth-microservices/users/models"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"net/mail"

	"golang.org/x/crypto/bcrypt"

	authpb "traileau-auth-microservices/traileau-gateway/proto/auth"
)

type ServerStruct struct {
	authpb.UnimplementedAuthServiceServer
	UserUseCase usecase.UserUsecase
}

func (s *ServerStruct) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	return &authpb.LoginResponse{Token: "access_token"}, nil
}

func (uc *ServerStruct) Signup(ctx context.Context, req *authpb.SignupRequest) (*authpb.SignupResponse, error) {
	// Initialize the validator

	var user model.User

	// Decode the request body to access the data like a json
	decoder := req
	if decoder.Username == "" {
		fmt.Printf("error %s", "username empty")
		return &authpb.SignupResponse{}, nil
	}
	// Use bcrypt to encrypt the password
	cryptedPassword, errorCryptingPassword := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if errorCryptingPassword != nil {
		fmt.Printf("error %s", "error crypting password")
		return &authpb.SignupResponse{}, nil
	}

	fmt.Println(req)

	validEmail, mailError := mail.ParseAddress(user.Email)

	if mailError != nil {
		fmt.Println("Email not valid")
		return &authpb.SignupResponse{}, nil
	}

	_, userExistingError := uc.UserUseCase.GetUser(ctx, &user.Email)

	if userExistingError == nil {
		fmt.Println("User already exist")
		return &authpb.SignupResponse{}, nil
	}

	newUser := model.User{
		Id:       primitive.NewObjectID(),
		Username: user.Username,
		Email:    validEmail.Address,
		Password: string(cryptedPassword),
	}

	// Insert the new user in db
	err := uc.UserUseCase.CreateUser(ctx, &newUser)
	if err != nil {
		return &authpb.SignupResponse{}, nil
	}

	return &authpb.SignupResponse{UserId: newUser.Username}, nil
}
