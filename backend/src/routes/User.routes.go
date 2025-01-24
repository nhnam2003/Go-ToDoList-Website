package routes

import (
	"github.com/DaiNef163/Go-ToDoList/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func RoutesUser(app *fiber.App) {
	api := app.Group("/api")

	// api.Get("/getuser", controllers.GetUser)
	api.Post("/getuser", controllers.GetUser)
}
