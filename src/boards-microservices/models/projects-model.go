package projects_model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" validate:"required" binding:"required"`
	Description string             `json:"description" validate:"required" binding:"required"`
	Author      string             `binding:"required"`
}
