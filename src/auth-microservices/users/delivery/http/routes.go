package http

import (
	"github.com/gin-gonic/gin"
)

func (uc UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/auth")
	userroute.POST("/signup", uc.SignUp)
	userroute.POST("signin", uc.SignIn)
	userroute.GET("/getall", uc.GetAll)
	userroute.GET("/get", uc.GetOne)
}
