package middleware

import (
	"strings"

	"stokeria/helper"

	"github.com/gofiber/fiber/v2"
)

// Protected memvalidasi Bearer token JWT dari header Authorization
func Protected(jwtSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "Missing Authorization header")
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid Authorization header format. Expected 'Bearer <token>'")
		}

		tokenString := parts[1]
		claims, err := helper.ValidateJWT(tokenString, jwtSecret)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired token: "+err.Error())
		}

		// Suntikkan data user ke dalam context Fiber
		c.Locals("user_id", claims.UserID)
		c.Locals("email", claims.Email)
		c.Locals("role", claims.Role)

		return c.Next()
	}
}

// RequireRole membatasi akses endpoint untuk role tertentu saja (misal: "admin")
func RequireRole(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		roleVal := c.Locals("role")
		if roleVal == nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized: user role not found")
		}

		userRole, ok := roleVal.(string)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized: invalid user role")
		}

		for _, allowed := range allowedRoles {
			if strings.EqualFold(userRole, allowed) {
				return c.Next()
			}
		}

		return fiber.NewError(fiber.StatusForbidden, "Forbidden: you do not have permission to access this resource")
	}
}
