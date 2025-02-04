package controllers

import (
    "time"
    "os"
    "github.com/DaiNef163/Go-ToDoList/src/config"
    "github.com/DaiNef163/Go-ToDoList/src/models"
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
    var data models.Account
    if err := c.BodyParser(&data); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

    // Validate input
    if data.Username == "" || data.Password == "" {
        return c.Status(400).JSON(fiber.Map{
            "message": "Username and password are required",
        })
    }

    // Hash password
    hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "message": "Error hashing password",
        })
    }

    // Create user
    user := models.Account{
        Username:  data.Username,
        Password:  string(hashPassword),
        Name:     data.Name,
        Age:      data.Age,
        Role:     data.Role,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    // Insert into database
    collection := config.GetCollection("Account")
    _, err = collection.InsertOne(c.Context(), user)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "message": "Error creating user",
        })
    }

    return c.Status(201).JSON(fiber.Map{
        "message": "Registration successful",
        "user": user,
    })
}

func Login(c *fiber.Ctx) error {
    var data models.Account
    if err := c.BodyParser(&data); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

    if data.Username == "" || data.Password == "" {
        return c.Status(400).JSON(fiber.Map{
            "message": "Username and password are required",
        })
    }

    // Find user in database
    var dataDB models.Account
    collection := config.GetCollection("Account")
    err := collection.FindOne(c.Context(), bson.M{"username": data.Username}).Decode(&dataDB)

    if err != nil {
        if err == mongo.ErrNoDocuments {
            return c.Status(400).JSON(fiber.Map{
                "message": "Invalid credentials",
            })
        }
        return c.Status(500).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

    // Compare passwords
    err = bcrypt.CompareHashAndPassword([]byte(dataDB.Password), []byte(data.Password))
    if err != nil {
        return c.Status(400).JSON(fiber.Map{
            "message": "Invalid credentials",
        })
    }

    // Create JWT token
    claims := jwt.MapClaims{
        "id":       dataDB.ID,
        "username": dataDB.Username,
        "role":     dataDB.Role,
        "exp":      time.Now().Add(time.Hour * 24).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        secret = "your-secret-key" // Only for development
    }
    
    tokenString, err := token.SignedString([]byte(secret))
    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "message": "Error generating token",
        })
    }

    return c.Status(200).JSON(fiber.Map{
        "message": "Login successful",
        "token": tokenString,
        "user": fiber.Map{
            "username": dataDB.Username,
            "role": dataDB.Role,
        },
    })
}