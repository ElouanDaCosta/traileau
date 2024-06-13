package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"traileau/configs"
	"traileau/users/delivery/http"
	domain "traileau/users/domain/usecase"
	"traileau/users/repository"
	"traileau/users/usecase"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

var (
	server      *gin.Engine
	us          domain.UserUsecase
	uc          http.UserController
	ctx         context.Context
	mongoClient *mongo.Client
)

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

	basePath := server.Group("/v1")
	uc.RegisterUserRoutes(basePath)

	dotenv := godotenv.Load()
	if dotenv != nil {
		log.Fatal("Error loading .env file")
	}
	server.SetTrustedProxies(nil)
	server.Run()

	log.Fatal(server.Run(":" + os.Getenv("AUTH_PORT")))
}
