package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	Name      string             `json:"name" bson:"name"`
	Age       int                `json:"age" bson:"age"`
	Role      string             `json:"role" bson:"role"` 
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

func (a *Account) SetCreatedAt() {
	if a.CreatedAt.IsZero() {
		a.CreatedAt = time.Now()
	}
}

func (a *Account) SetUpdatedAt() {
	a.UpdatedAt = time.Now()
}
