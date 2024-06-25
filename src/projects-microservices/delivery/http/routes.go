package http

import (
	"traileau-projects-microservices/delivery/middleware"

	"github.com/gin-gonic/gin"
)

func (pc ProjectController) RegisterProjectRoutes(rg *gin.RouterGroup) {
	projectroute := rg.Group("project")
	projectroute.GET("/", middleware.Authenticate(), pc.GetAll)
	projectroute.POST("/", pc.CreateProject)
}
