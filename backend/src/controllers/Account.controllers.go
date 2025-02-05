package controllers

import (
	"errors"
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

func CreateJWTToken(username, role, userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["role"] = role
	claims["userId"] = userID

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET is not set in environment variables")
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Register(c *fiber.Ctx) error {
	var data models.Account
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Dữ liệu không hợp lệ",
			"error":   err.Error(),
		})
	}

	if data.Username == "" || data.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Username và password không được để trống",
		})
	}

	// Kiểm tra username đã tồn tại
	collection := config.GetCollection("Account")
	existingUser := models.Account{}
	err := collection.FindOne(c.Context(), bson.M{"username": data.Username}).Decode(&existingUser)
	if err == nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Username đã tồn tại",
		})
	} else if !errors.Is(err, mongo.ErrNoDocuments) {
		return c.Status(500).JSON(fiber.Map{
			"message": "Lỗi server",
			"error":   err.Error(),
		})
	}

	hashPwd, err := hashpassword.HashPassword(data.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Lỗi khi mã hóa mật khẩu",
			"error":   err.Error(),
		})
	}

	user := models.Account{
		Username:  data.Username,
		Password:  hashPwd,
		Name:      data.Name,
		Age:       data.Age,
		Role:      data.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if _, err := collection.InsertOne(c.Context(), user); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Lỗi khi tạo tài khoản",
			"error":   err.Error(),
		})
	}

	// Ẩn mật khẩu trước khi trả về
	user.Password = ""
	return c.Status(201).JSON(fiber.Map{
		"message": "Đăng ký thành công",
		"user":    user,
	})
}

func Login(c *fiber.Ctx) error {
	var data models.Account
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Dữ liệu không hợp lệ",
			"error":   err.Error(),
		})
	}

	if data.Username == "" || data.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Username và password không được để trống",
		})
	}

	var userDB models.Account
	collection := config.GetCollection("Account")
	err := collection.FindOne(c.Context(), bson.M{"username": data.Username}).Decode(&userDB)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(401).JSON(fiber.Map{
				"message": "Username hoặc password không chính xác",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"message": "Lỗi server",
			"error":   err.Error(),
		})
	}

	// Kiểm tra mật khẩu
	if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(data.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Username hoặc password không chính xác",
		})
	}

	// Tạo JWT token
	token, err := CreateJWTToken(userDB.Username, userDB.Role, userDB.ID.Hex()) // Sử dụng .Hex() để chuyển ObjectID thành string
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Lỗi khi tạo token",
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Đăng nhập thành công",
		"token":   token,
		"user": fiber.Map{
			"username": userDB.Username,
			"role":     userDB.Role,
			"userId":   userDB.ID.Hex(), // Sử dụng .Hex() ở đây cũng vậy
		},
	})
}

