package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"traileau-list-microservices/configs"
	"traileau-list-microservices/delivery/http"
	domain "traileau-list-microservices/domain/usecase"
	"traileau-list-microservices/repository"
	"traileau-list-microservices/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server         *gin.Engine
	mongoClient    *mongo.Client
	ctx            context.Context
	listusecase    domain.ListUseCase
	listcontroller http.ListController
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

	listrepository := repository.NewProjectRepository(mongoCon)
	listusecase = usecase.NewProjectUsecase(listrepository, ctx)
	listcontroller = http.New(listusecase)

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

	basePath := server.Group("v" + os.Getenv("LIST_API_VERSION"))
	listcontroller.RegisterListRoutes(basePath)
	server.SetTrustedProxies(nil)
	server.Run(":" + os.Getenv("LIST_PORT"))

	log.Fatal(server.Run(":" + os.Getenv("LIST_PORT")))
}
