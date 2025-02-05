package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ToDo struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID    string             `json:"userId" bson:"userId"`
	Title     string             `json:"title" bson:"title"`
	Complete  bool               `json:"complete" bson:"complete"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
