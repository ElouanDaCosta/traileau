package main

import (
	"log"
	"net/http"
	"os"
	"traileau/configs"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	dotenv := godotenv.Load()
	if dotenv != nil {
		log.Fatal("Error loading .env file")
	}
	configs.Connect()
	r.SetTrustedProxies(nil)
	err := r.Run(":" + os.Getenv("MAIN_PORT"))
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
	r.Run()
}

// func SetupAppRouter() *gin.Engine {

// 	db := configs.Connection()

// 	router := gin.Default()

// 	gin.SetMode(gin.TestMode)

// 	api := router.Group("api/v1")

// 	routes.InitAuthRoutes(db, api)

// 	return router
// }
