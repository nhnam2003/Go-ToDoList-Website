package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DaiNef163/Go-ToDoList/src/config"
	"github.com/DaiNef163/Go-ToDoList/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("❌ Lỗi khi tải file .env")
	}

	// Kết nối MongoDB
	config.MongoDB()


	// Đóng kết nối khi ứng dụng dừng
	defer config.CloseDB()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server is running!")
	})

	// Setup routes
	routes.RoutesAccount(app)
	routes.RoutesTodo(app)
	routes.RoutesUser(app)

	// Lấy PORT từ biến môi trường
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// if os.Getenv("ENV") == "production" {
	// 	app.Static("/", "./client/dist")
	// }

	fmt.Println("Server đang chạy tại PORT:", port)
	log.Fatal(app.Listen("0.0.0.0:" + port))
}
