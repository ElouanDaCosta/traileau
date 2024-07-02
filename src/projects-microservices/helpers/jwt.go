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
	email     string
	sub       int
	ExpiresAt int
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func ExtractUnverifiedClaims(tokenString string) (string, error) {
	var email string
	parsedToken := strings.Split(tokenString, "Bearer ")

	token, _, err := new(jwt.Parser).ParseUnverified(parsedToken[1], jwt.MapClaims{})

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		email = fmt.Sprintln(claims["email"])
	}

	if email == "" {
		return "", fmt.Errorf("invalid token payload")
	}
	return email, nil
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

func GetTokenData(ctx *gin.Context) (string, error) {
	token := ExtractToken(ctx)

	if token == "nul" {
		return "nul", fmt.Errorf("error getting the token")
	}

	emailFromToken, err := ExtractUnverifiedClaims(token)

	if err != nil {
		return "Email: nul", err
	}

	return emailFromToken, err
}
