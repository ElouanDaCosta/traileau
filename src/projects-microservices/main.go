package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"traileau-projects-microservices/configs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server      *gin.Engine
	mongoClient *mongo.Client
	ctx         context.Context
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

	// ur := repository.NewUserRepository(mongoCon)
	// us = usecase.NewUserUsecase(ur, ctx)
	// uc = http.New(us)

	server = gin.Default()
}

func main() {
	dotenv := godotenv.Load()
	if dotenv != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Projects pong",
		})
	})

	// basePath := server.Group("v" + os.Getenv("AUTH_API_VERSION"))
	server.SetTrustedProxies(nil)
	server.Run(":" + os.Getenv("PROJECTS_PORT"))

	log.Fatal(server.Run(":" + os.Getenv("PROJECTS_PORT")))
}
