package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"traileau/configs"
	userController "traileau/users/controllers"
	domain "traileau/users/domain/repository"
	usecase "traileau/users/usecase"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

var (
	engine        *gin.Engine
	userService   domain.UserRepositoryInterface
	usecontroller userController.UserController
	ctx           context.Context
	mongoClient   *mongo.Client
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
	r := gin.Default()
	dotenv := godotenv.Load()
	if dotenv != nil {
		log.Fatal("Error loading .env file")
	}

	ur :=

		r.SetTrustedProxies(nil)
	err := r.Run(":" + os.Getenv("MAIN_PORT"))
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
	r.Run()
}

func main() {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

// func SetupAppRouter() *gin.Engine {

// 	db := configs.Connection()

// 	router := gin.Default()

// 	gin.SetMode(gin.TestMode)

// 	api := router.Group("api/v1")

// 	routes.InitAuthRoutes(db, api)

// 	return router
// }
