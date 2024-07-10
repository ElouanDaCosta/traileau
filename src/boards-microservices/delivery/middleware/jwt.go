package middleware

import (
	helper "boards-projects-microservices/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := helper.ExtractToken(c)

		_, err := helper.ExtractUnverifiedClaims(token)

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"Error": "authenticate with user token"})
			c.Abort()
			return
		}
	}
}
