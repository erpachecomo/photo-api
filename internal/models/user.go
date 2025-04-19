package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID    bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name  string        `json:"name" binding:"required"`
	Email string        `json:"email" binding:"required"`
}
