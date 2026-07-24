package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func Protect() fiber.Handler {
	return func(req fiber.Ctx) error {

		// Getting Auth Header
		authHeader := req.Get("Authorization")
		if authHeader == "" {
			return req.Status(401).JSON(fiber.Map{"status": "unauthorized"})
		}
		fmt.Println("Before Trimming: ", authHeader)
		authHeader = strings.TrimPrefix(authHeader, "Bearer ")

		// Checking JWT
		fmt.Println("After Trim: ", authHeader)

		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			return err
		}
		if !token.Valid {
			return req.Status(401).JSON(fiber.Map{
				"status": "invalid token",
			})
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return req.Status(500).JSON(fiber.Map{"status": "unexpected error while getting claims"})
		}
		fmt.Println(claims["user_id"])
		fmt.Println(claims["email"])

		req.Locals("user_id", claims["user_id"])
		req.Locals("email", claims["email"])

		return req.Next()
	}
}
