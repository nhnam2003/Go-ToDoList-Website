package routes

import (
	"github.com/DaiNef163/Go-ToDoList/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func RoutesItem(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/pv", controllers.Login)


}
