package middlewares

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// Auth middleware kiểm tra và xác thực token trong header Authorization
func Auth(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Không có token",
		})
	}

	// Lấy token từ Authorization header (format: Bearer <token>)
	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Token không hợp lệ",
		})
	}
	tokenString = parts[1]

	// Giải mã token với secret từ biến môi trường
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Sử dụng JWT_SECRET từ biến môi trường
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		if err.Error() == "Token is expired" {
			return c.Status(401).JSON(fiber.Map{
				"message": "Token hết hạn",
			})
		}
		return c.Status(401).JSON(fiber.Map{
			"message": "Xác thực người dùng thất bại",
			"error":   err.Error(),
		})
	}

	// Kiểm tra tính hợp lệ của token
	if !token.Valid {
		return c.Status(401).JSON(fiber.Map{
			"message": "Token không hợp lệ",
		})
	}

	// Lưu thông tin từ token vào context để sử dụng trong các handler sau
	claims := token.Claims.(jwt.MapClaims)

	// Đảm bảo các claim như userId tồn tại trong token
	userId, ok := claims["userId"].(string)
	if !ok {
		return c.Status(401).JSON(fiber.Map{
			"message": "Không tìm thấy thông tin người dùng trong token",
		})
	}

	// Lưu thông tin người dùng vào context
	c.Locals("user", claims)
	c.Locals("userId", userId)

	return c.Next()
}

// RoleGuard kiểm tra quyền của người dùng dựa trên vai trò
func RoleGuard(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(jwt.MapClaims)
		userRole, ok := user["role"].(string)
		if !ok {
			return c.Status(403).JSON(fiber.Map{
				"message": "Không tìm thấy vai trò người dùng",
			})
		}

		// Kiểm tra xem vai trò người dùng có hợp lệ hay không
		for _, role := range roles {
			if role == userRole {
				return c.Next()
			}
		}

		return c.Status(403).JSON(fiber.Map{
			"message": "Bạn không có quyền truy cập tài nguyên này",
		})
	}
}
