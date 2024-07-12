package http

import "github.com/gin-gonic/gin"

func (lc ListController) RegisterListRoutes(rg *gin.RouterGroup) {
	listroute := rg.Group("lists")
	listroute.GET("/", lc.GetAll)
	listroute.POST("/", lc.CreateList)
}
