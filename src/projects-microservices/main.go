package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"traileau-projects-microservices/configs"
	"traileau-projects-microservices/delivery/http"
	domain "traileau-projects-microservices/domain/usecase"
	"traileau-projects-microservices/repository"
	"traileau-projects-microservices/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server      *gin.Engine
	mongoClient *mongo.Client
	ctx         context.Context
	ps          domain.ProjectUseCase
	pc          http.ProjectController
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

	pr := repository.NewProjectRepository(mongoCon)
	ps = usecase.NewProjectUsecase(pr, ctx)
	pc = http.New(ps)

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
	pc.RegisterProjectRoutes(basePath)
	server.SetTrustedProxies(nil)
	server.Run(":" + os.Getenv("PROJECTS_PORT"))

	log.Fatal(server.Run(":" + os.Getenv("PROJECTS_PORT")))
}
