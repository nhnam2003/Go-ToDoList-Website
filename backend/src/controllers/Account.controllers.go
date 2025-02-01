package controllers

import (
	hashpassword "github.com/DaiNef163/Go-ToDoList/src/service/hashPassword"
	"github.com/gofiber/fiber/v2"
)

func CheckHassPassword(c *fiber.Ctx) error{
	pw1 := "abc123"
	hash,_ := hashpassword.HashPassword(pw1)
	return c.JSON(hash)

}