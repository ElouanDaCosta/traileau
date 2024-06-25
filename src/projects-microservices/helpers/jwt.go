package helper

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type SignedDetails struct {
	Email     string
	Username  string
	User_type string
	ExpiresAt int
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func ExtractUnverifiedClaims(tokenString string) (string, error) {
	var name string
	parsedToken := strings.Split(tokenString, "Bearer ")
	token, _, err := new(jwt.Parser).ParseUnverified(parsedToken[1], jwt.MapClaims{})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		name = fmt.Sprintln(claims["sub"])
	}

	if name == "" {
		return "", fmt.Errorf("invalid token payload")
	}
	return name, nil
}

func ExtractToken(ctx *gin.Context) string {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Token not found"})
		ctx.Abort()
		return "nul"
	}

	return token
}
