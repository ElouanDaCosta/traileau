package helper

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type SignedDetails struct {
	Email     string
	Username  string
	User_type string
	ExpiresAt int
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func ExtractToken(ctx *gin.Context) (claims *SignedDetails, msg string) {
	token, err := ctx.Request.Header["Authorization"]

	if !err {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Token not found"})
		ctx.Abort()
		return
	}

	fmt.Println(token)

	// if err != true {
	// 	msg = err.Error()
	// 	return
	// }

	// claims, ok := token.Claims.(*SignedDetails)
	// if !ok {
	// 	msg = fmt.Sprintf("the token is invalid")
	// 	msg = err.Error()
	// 	return
	// }

	// if claims.ExpiresAt < time.Now().Local().Unix() {
	// 	msg = fmt.Sprintf("token is expired")
	// 	msg = err.Error()
	// 	return
	// }
	return claims, msg
}
