package controllers

import (
	"context"

	"github.com/DaiNef163/Go-ToDoList/src/config"
	"github.com/DaiNef163/Go-ToDoList/src/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	// Kiểm tra kết nối MongoDB và lấy đúng collection
	insertResult, err := config.GetCollection("TodoGo", "user").InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	user.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.JSON(user)
}
