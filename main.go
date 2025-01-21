package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type ToDo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}


func main() {
	err:= godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")

	fmt.Println("Hello World")
	app := fiber.New()
	todos := []ToDo{}

	app.Get("/", func(c *fiber.Ctx) error {
		todo := &ToDo{}
		fmt.Println("log all todo")
		fmt.Println(&todo)

		return c.Status(200).JSON(todos)
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &ToDo{}
		if err := c.BodyParser(todo); err != nil {
			return err
		}
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{
				"message": "Body is required",
			})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})

	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{
			"message": "Todo not found",
		})
	})

	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				//cắt chuỗi sau đó nối đầu đuôi
				// [1,2,3,4,5]
				//  i = 2
				// todos[:i] , sẽ là [1,2]
				//  todos[i+1:] , sẽ là [4,5]
				return c.Status(200).JSON(fiber.Map{
					"message": "Todo deleted",
				})
			}
		}
		return c.Status(404).JSON(fiber.Map{
			"message": "Todo not found",
		})
	})

	// arr := [5]int{1, 2, 3, 4, 5}
	// fmt.Println("arr[:2]", arr[:2])
	// fmt.Println("arr[2+1:]", arr[2+1:])

	log.Fatal(app.Listen(":"+PORT))

}
