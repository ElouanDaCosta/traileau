package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SplitObjectID(objectID primitive.ObjectID) string {
	hexString := objectID.Hex()

	part1 := hexString[:8]
	part2 := hexString[8:16]
	part3 := hexString[16:]

	objectId := part1 + part2 + part3

	return objectId
}
