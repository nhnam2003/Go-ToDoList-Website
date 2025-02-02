package controllers

import (
	"fmt"

	"github.com/DaiNef163/Go-ToDoList/src/config"
	"github.com/DaiNef163/Go-ToDoList/src/models"
	hashpassword "github.com/DaiNef163/Go-ToDoList/src/service/hashPassword"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func CheckHassPassword(c *fiber.Ctx) error {
	pw1 := "abc123"
	hash, _ := hashpassword.HashPassword(pw1)
	return c.JSON(hash)

}

func Register(c *fiber.Ctx) error {

	var data models.Account
	err1 := c.BodyParser(&data)
	if err1 != nil {
		return err1
	}
	var hashPassword, _ = hashpassword.HashPassword(data.Password)

	user := models.Account{
		Username: data.Username,
		Password: hashPassword,
		Name:     data.Name,
		Age:      data.Age,
		Role:     data.Role,
	}

	collection := config.GetCollection("Account")
	_, err2 := collection.InsertOne(c.Context(), user)
	if err2 != nil {
		return err2
	}
	return c.JSON(user)

}

func Login(c *fiber.Ctx) error {
	var data models.Account

	err1 := c.BodyParser(&data)
	if err1 != nil {
		return c.Status(400).JSON(err1.Error())
	}
	fmt.Println("&data", &data)
	if data.Username == "" || data.Password == "" {
		return c.Status(400).JSON("Username or Password is empty")
	}

	var dataDB models.Account
	var collection = config.GetCollection("Account")
	err := collection.FindOne(c.Context(), bson.M{"username": data.Username}).Decode(&dataDB)
	// err := collection.FindOne(c.Context(), bson.M{
	// 	"username": data.Username,
	// }).Decode(&existingUser)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(400).JSON("Username or Password is incorrect")
		}
		return c.Status(400).JSON(err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(dataDB.Password), []byte(data.Password))
	if err != nil {
		return c.Status(400).JSON("Username or Password is incorrect")
	}

	return c.Status(200).JSON("Login success ")
}
