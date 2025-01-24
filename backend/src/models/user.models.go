package models

import (
	"github.com/DaiNef163/Go-ToDoList/src/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Age      int                `json:"age" bson:"age"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
	Role     []string           `json:"role" bson:"role"`
}

var userCollection *mongo.Collection

// Khởi tạo userCollection
func InitUserCollection() {
	// Đảm bảo kết nối MongoDB đã thành công trước khi lấy collection
	userCollection = config.GetCollection("userDB", "TodoGo")
}
