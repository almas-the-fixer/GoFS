package routes

import (
	"gofs/handlers"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
)

func RegisterRoutes(app *fiber.App, conn *pgx.Conn) {
	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "healthy"})
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// User CRUD Need To Refactor Later as Admin / User Endpoints...
	app.Get("/user/:id", handlers.GetUser(conn))
	app.Get("/users", handlers.GetUsers(conn))
	app.Delete("/users/:id", handlers.DeleteUser(conn))
	app.Post("/users", handlers.CreateUser(conn))
}
