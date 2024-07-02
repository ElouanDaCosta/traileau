package http

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var dotenv = godotenv.Load()

var (
	auth_baseUrl       string = os.Getenv("AUTH_API_URL")
	auth_version       string = os.Getenv("AUTH_API_VERSION")
	auth_port          string = os.Getenv("AUTH_PORT")
	auth_route         string = os.Getenv("AUTH_ROUTE")
	authUrl            string = auth_baseUrl + auth_port + "/v" + auth_version + "/" + auth_route + "/"
	getoneuserEndpoint string = os.Getenv("AUTH_GET_ONE_USER_ENDPOINT")
)

func init() {
	if dotenv != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetUser(email string) (*http.Response, error) {

	temp := strings.Split(email, "\n")

	resp, err := http.Get(authUrl + getoneuserEndpoint + temp[0])

	if err != nil {
		fmt.Println(err)
		return resp, err
	}

	fmt.Println("Request send to the authentication microservice at:", auth_baseUrl+auth_port)

	return resp, err
}
