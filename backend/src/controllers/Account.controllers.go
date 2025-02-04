package controllers

import (
	"fmt"
	"os"
	"time"

	"github.com/DaiNef163/Go-ToDoList/src/config"
	"github.com/DaiNef163/Go-ToDoList/src/models"
	hashpassword "github.com/DaiNef163/Go-ToDoList/src/service/hashPassword"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
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
		Username:  data.Username,
		Password:  hashPassword,
		Name:      data.Name,
		Age:       data.Age,
		Role:      data.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = dataDB.Username
	claims["role"] = dataDB.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	serect := os.Getenv("JWT_SECRET")

	if serect == "" {
		return c.Status(500).JSON("JWT_SECRET is empty")
	}

	tokenString, err := token.SignedString([]byte(serect))
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Login successful",
		"token":   tokenString,
		"user": fiber.Map{
			"username": dataDB.Username,
			"role":     dataDB.Role,
		},
	})
}
