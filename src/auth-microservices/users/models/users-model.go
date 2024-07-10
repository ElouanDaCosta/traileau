package users_model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string             `json:"username" validate:"required" binding:"required"`
	Email    string             `json:"email" validate:"required" binding:"required"`
	Password string             `json:"password" validate:"required" binding:"required"`
}
