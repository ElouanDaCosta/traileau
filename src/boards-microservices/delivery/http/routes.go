package http

import (
	"boards-projects-microservices/delivery/middleware"

	"github.com/gin-gonic/gin"
)

func (pc ProjectController) RegisterProjectRoutes(rg *gin.RouterGroup) {
	projectroute := rg.Group("project")
	projectroute.GET("/", middleware.Authenticate(), pc.GetAll)
	projectroute.POST("/", middleware.Authenticate(), pc.CreateProject)
}
