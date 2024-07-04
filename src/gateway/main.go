package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	projectpb "traileau-gateway/traileau-gateway/proto/project"

	authpb "traileau-gateway/traileau-gateway/proto/auth"
)

func main() {
	r := gin.Default()

	// Connection to the Grpc services
	authConn, err := grpc.NewClient("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to AuthService: %v", err)
	}
	defer authConn.Close()
	authClient := authpb.NewAuthServiceClient(authConn)

	projectConn, err := grpc.NewClient("tasks-service:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to TaskService: %v", err)
	}
	defer projectConn.Close()
	projectClient := projectpb.NewProjectServiceClient(projectConn)

	// Http routes
	r.POST("/v1/auth/signin", func(c *gin.Context) {
		var req authpb.LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := authClient.Login(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	})

	r.POST("/v1/auth/signup", func(c *gin.Context) {
		var req authpb.SignupRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := authClient.Signup(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	})

	r.POST("/projects", func(c *gin.Context) {
		var req projectpb.CreateProjectRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := projectClient.CreateProject(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	})

	// Run the Gateway
	r.Run(":8080")
}
