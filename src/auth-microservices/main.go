package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"traileau-auth-microservices/configs"
	"traileau-auth-microservices/users/delivery/http"
	domain "traileau-auth-microservices/users/domain/usecase"
	"traileau-auth-microservices/users/repository"
	"traileau-auth-microservices/users/usecase"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"

	authpb "traileau-auth-microservices/traileau-gateway/proto/auth"
)

var (
	server      *gin.Engine
	us          domain.UserUsecase
	uc          http.UserController
	ctx         context.Context
	mongoClient *mongo.Client
)

type serverStruct struct {
	authpb.UnimplementedAuthServiceServer
}

func (s *serverStruct) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	// Implémentation de la fonction Login
	return &authpb.LoginResponse{Token: "example_token"}, nil
}

func init() {
	ctx = context.TODO()

	// mongo
	mongoCon, err := configs.Connect()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("mongo connection established")

	ur := repository.NewUserRepository(mongoCon)
	us = usecase.NewUserUsecase(ur, ctx)
	uc = http.New(us)

	server = gin.Default()
}

func main() {
	defer func(mongoClient *mongo.Client, ctx context.Context) {
		err := mongoClient.Disconnect(ctx)
		if err != nil {
			log.Println(err)
		}
	}(mongoClient, ctx)

	dotenv := godotenv.Load()
	if dotenv != nil {
		log.Fatal("Error loading .env file")
	}
	basePath := server.Group("v" + os.Getenv("AUTH_API_VERSION"))
	uc.RegisterUserRoutes(basePath)
	server.SetTrustedProxies(nil)
	// server.Run(":" + os.Getenv("AUTH_PORT"))

	// log.Fatal(server.Run(":" + os.Getenv("AUTH_PORT")))

	grpcServer, grpcErr := net.Listen("tcp", ":8081")

	if grpcErr != nil {
		fmt.Println(grpcErr)
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &serverStruct{})
	log.Println("Starting AuthService server on port 8081...")
	if grpcErr := s.Serve(grpcServer); grpcErr != nil {
		log.Fatalf("Failed to serve: %v", grpcErr)
	}
}
