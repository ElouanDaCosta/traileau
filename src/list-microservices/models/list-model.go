package list_model

import "go.mongodb.org/mongo-driver/bson/primitive"

type List struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" validate:"required" binding:"required"`
	Position int                `json:"position" binding:"required"`
	Boards   string             `json:"boards" binding:"required"`
}
