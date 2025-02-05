package controllers

import (
	"context"
	"fmt"
	"log"

	// "fmt"
	// "go.mongodb.org/mongo-driver/bson"

	"github.com/DaiNef163/Go-ToDoList/src/config"
	"github.com/DaiNef163/Go-ToDoList/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"

	// "github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTodos(c *fiber.Ctx) error {
	// Lấy thông tin người dùng từ context (được lưu bởi middleware Auth)
	user := c.Locals("user").(jwt.MapClaims)
	userID := user["userId"].(string)

	// Lấy collection "todos" từ MongoDB
	todoCollection := config.GetCollection("todo")

	// Tìm các todo của người dùng
	var todos []models.ToDo
	filter := bson.M{"userId": userID}
	cursor, err := todoCollection.Find(context.TODO(), filter)
	if err != nil {
		log.Printf("❌ Lỗi khi tìm todos: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Lỗi khi lấy dữ liệu todo",
		})
	}
	defer cursor.Close(context.TODO())

	// Duyệt qua các kết quả và thêm vào danh sách todos
	for cursor.Next(context.TODO()) {
		var todo models.ToDo
		if err := cursor.Decode(&todo); err != nil {
			log.Printf("❌ Lỗi khi giải mã todo: %v", err)
			return c.Status(500).JSON(fiber.Map{
				"message": "Lỗi khi giải mã dữ liệu todo",
			})
		}
		todos = append(todos, todo)
	}

	// Kiểm tra lỗi sau khi duyệt hết cursor
	if err := cursor.Err(); err != nil {
		log.Printf("❌ Lỗi khi duyệt todos: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Lỗi khi duyệt dữ liệu todo",
		})
	}

	// Trả về danh sách todos
	return c.JSON(fiber.Map{
		"todo": todos,
	})
}

func CreateTodo(c *fiber.Ctx) error {
	// Lấy thông tin người dùng từ context
	user := c.Locals("user").(jwt.MapClaims)
	userID := user["userId"].(string)

	// Khai báo một đối tượng Todo mới từ body request
	todo := new(models.ToDo)

	log.Printf("Dữ liệu nhận được: %+v", todo)
	// Parse dữ liệu từ body
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Dữ liệu không hợp lệ",
			"error":   err.Error(),
		})
	}

	// Set userID cho todo
	todo.UserID = userID
	todo.Complete = false // mặc định là chưa hoàn thành

	// Lấy collection "todos" từ MongoDB
	collection := config.GetCollection("todo")

	// Thêm todo vào MongoDB
	result, err := collection.InsertOne(c.Context(), todo)
	if err != nil {
		log.Println("❌ Lỗi khi thêm todo:", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Lỗi khi thêm todo",
			"error":   err.Error(),
		})
	}

	// Trả về kết quả
	return c.Status(201).JSON(fiber.Map{
		"message": "Todo đã được tạo thành công",
		"todo":    result,
	})
}

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": ObjectID}
	update := bson.M{"$set": bson.M{"completed": true}}
	result, err := config.GetCollection("todo").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	fmt.Printf("ID: %s\n", id)
	return c.JSON(fiber.Map{"message": "Todo updated"})
}

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
