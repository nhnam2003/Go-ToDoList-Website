package controllers

import (
	"context"
	"fmt"
	"log"
	"time"

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
	user := c.Locals("user").(jwt.MapClaims)
	userID := user["userId"].(string)

	todoCollection := config.GetCollection("todo")

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

	return c.JSON(fiber.Map{
		"todo": todos,
	})
}

func CreateTodo(c *fiber.Ctx) error {
	user := c.Locals("user").(jwt.MapClaims)
	userID := user["userId"].(string)

	todo := new(models.ToDo)

	log.Printf("Dữ liệu nhận được: %+v", todo)
	// Parse dữ liệu từ body
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Dữ liệu không hợp lệ",
			"error":   err.Error(),
		})
	}

	todo.UserID = userID
	todo.Complete = false 
	todo.CreatedAt = time.Now()


	collection := config.GetCollection("todo")


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
	user := c.Locals("user").(jwt.MapClaims)
	userID := user["userId"].(string)

	id := c.Params("id")
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// Lọc todo theo _id
	filter := bson.M{"_id": ObjectID, "userId": userID} // Đảm bảo chỉ cập nhật của người dùng hiện tại
	update := bson.M{
		"$set": bson.M{
			"complete":  true,                // Cập nhật trường complete
			"updated_at": time.Now(),         // Cập nhật thời gian update
		},
	}

	// Thực hiện cập nhật
	result, err := config.GetCollection("todo").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	// Kiểm tra nếu không có bản ghi nào được tìm thấy và cập nhật
	if result.MatchedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found or unauthorized",
		})
	}

	// In ra ID đã cập nhật (cho debug)
	fmt.Printf("Updated Todo ID: %s\n", id)

	return c.JSON(fiber.Map{
		"message": "Todo updated successfully",
	})
}



func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": ObjectID}
	result, err := config.GetCollection("todo").DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}
	return c.JSON(fiber.Map{"message": "Todo deleted"})
}
