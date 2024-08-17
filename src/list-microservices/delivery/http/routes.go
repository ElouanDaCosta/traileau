package http

import (
	"traileau-list-microservices/delivery/middleware"

	"github.com/gin-gonic/gin"
)

func (lc ListController) RegisterListRoutes(rg *gin.RouterGroup) {
	listroute := rg.Group("lists")
	listroute.GET("/", middleware.Authenticate(), lc.GetAll)
	listroute.POST("/", middleware.Authenticate(), lc.CreateList)
	listroute.DELETE("/", middleware.Authenticate(), lc.DeleteList)
}
