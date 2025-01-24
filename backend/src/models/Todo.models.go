package models

import (
	"github.com/DaiNef163/Go-ToDoList/src/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ToDo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

var todoCollection *mongo.Collection

func InitTodoCollection() {
	todoCollection = config.GetCollection("todoDB", "TodoGo")
}
