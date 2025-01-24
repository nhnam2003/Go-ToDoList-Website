package controllers

import (
	"context"

	// "fmt"
	// "go.mongodb.org/mongo-driver/bson"

	"github.com/DaiNef163/Go-ToDoList/src/config"
	"github.com/DaiNef163/Go-ToDoList/src/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTodos(c *fiber.Ctx) error {
	var todos []models.ToDo

	cursor, err := config.GetCollection("todo").Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo models.ToDo
		if err := cursor.Decode(&todo); err != nil {
			return err
		}
		todos = append(todos, todo)
	}
	return c.JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
	todo := new(models.ToDo)
	if err := c.BodyParser(todo); err != nil {
		return err
	}

	// Kiểm tra kết nối MongoDB và lấy đúng collection
	insertResult, err := config.GetCollection("todo").InsertOne(context.TODO(), todo)
	if err != nil {
		return err
	}

	todo.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.JSON(todo)
}

// func UpdateTodo(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	ObjectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return err
// 	}

// 	filter := bson.M{"_id": ObjectID}
// 	update := bson.M{"$set": bson.M{"completed": true}}
// 	result, err := config.GetDB().UpdateOne(context.Background(), filter, update)
// 	if err != nil {
// 		return err
// 	}

// 	if result.MatchedCount == 0 {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 			"error": "Todo not found",
// 		})
// 	}

// 	fmt.Printf("ID: %s\n", id)
// 	return c.JSON(fiber.Map{"message": "Todo updated"})
// }

// func DeleteTodo(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	ObjectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return err
// 	}

// 	filter := bson.M{"_id": ObjectID}
// 	result, err := config.GetDB().DeleteOne(context.Background(), filter)
// 	if err != nil {
// 		return err
// 	}

// 	if result.DeletedCount == 0 {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 			"error": "Todo not found",
// 		})
// 	}
// 	return c.JSON(fiber.Map{"message": "Todo deleted"})
// }
