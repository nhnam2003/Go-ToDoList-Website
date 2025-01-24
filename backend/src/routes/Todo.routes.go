package routes

import (
	"github.com/DaiNef163/Go-ToDoList/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func RoutesTodo(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/gettodos", controllers.GetTodos)
	api.Post("/createtodos", controllers.CreateTodo)
	// api.Patch("/updatetodos/:id", controllers.UpdateTodo)
	// api.Delete("/deletetodos/:id", controllers.DeleteTodo)
}
