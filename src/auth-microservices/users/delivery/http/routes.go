package http

import (
	"github.com/gin-gonic/gin"
)

func (uc UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/auth")
	userroute.POST("/signup", uc.Register)
	userroute.GET("/get", uc.GetAll)
}
