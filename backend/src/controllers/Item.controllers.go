package controllers

import (
	"context"
	"encoding/json"

	"github.com/DaiNef163/Go-ToDoList/src/config"
	"github.com/DaiNef163/Go-ToDoList/src/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)	

// CountItems - Đếm số lượng item giống nhau
func CountItems(c *fiber.Ctx) error {
	collection := config.GetCollection("items")

	// Lấy tất cả item từ MongoDB
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch items"})
	}
	defer cursor.Close(context.Background())

	// Lưu danh sách item vào slice
	var items []models.Item
	if err := cursor.All(context.Background(), &items); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse items"})
	}

	// Đếm số lượng xuất hiện của từng item
	counts := make(map[string]int)
	for _, item := range items {
		key, _ := json.Marshal(item) // Chuyển item thành JSON string
		counts[string(key)]++
	}

	// Trả về kết quả đếm số lượng
	return c.JSON(counts)
}
