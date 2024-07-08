package configs

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Database, error) {
	dotenv := godotenv.Load()
	if dotenv != nil {
		log.Fatal("Error loading .env file")
	}
	var cred options.Credential = options.Credential{
		Username: os.Getenv("MONGO_DB_USER"),
		Password: os.Getenv("MONGO_DB_PASS"),
	}
	connPattern := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(connPattern).SetAuth(cred)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	return client.Database("mydb"), err
}
