package controllers

import (
	"github.com/gin-gonic/gin"
)

func (uc UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/create", uc.CreateUser)
}
