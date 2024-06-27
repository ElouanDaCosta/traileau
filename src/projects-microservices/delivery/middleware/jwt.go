package middleware

import (
	"net/http"
	helper "traileau-projects-microservices/helpers"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := helper.ExtractToken(c)

		_, err := helper.ExtractUnverifiedClaims(token)

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"message": "zizi"})
			c.Abort()
			return
		}

		// c.Set("email", claims.Email)
		// c.Set("first_name", claims.Username)
		// c.Set("user_type", claims.User_type)
		// c.Next()
	}
}
