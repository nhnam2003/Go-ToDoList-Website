package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Auth(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "No token provided",
		})
	}

	// Xóa "Bearer " nếu có
	//   if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
	// 	tokenString = tokenString[7:]
	// }
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRECT")), nil
	})

	if err != nil && !token.Valid {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}
	// Lưu claims vào context
	claim := token.Claims.(jwt.MapClaims)
	c.Locals("user", claim)

	return c.Next()
}

func RoleGuard(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(jwt.MapClaims)
		userRole := user["role"].(string)

		for _, role := range roles {
			if role == userRole {
				return c.Next()
			}
		}
		return c.Status(403).JSON(fiber.Map{
			"message": "You don't have permission to access this resource",
		})
	}
}
