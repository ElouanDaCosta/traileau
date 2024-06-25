package middleware

import (
	"fmt"
	"net/http"
	helper "traileau-projects-microservices/helpers"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization Header Provided")})
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken, c)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("first_name", claims.Username)
		c.Set("user_type", claims.User_type)
		c.Next()
	}
}
