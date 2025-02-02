package routes

import (
	"github.com/DaiNef163/Go-ToDoList/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func RoutesAccount(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/account", controllers.CheckHassPassword)
	api.Post("/login", controllers.Login)
	api.Post("/register", controllers.Register)

}
