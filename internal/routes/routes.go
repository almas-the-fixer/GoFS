package routes

import (
	"gofs/handlers"
	"gofs/internal/middleware"

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
	app.Get("/users", middleware.Protect(), handlers.GetUsers(conn))
	app.Delete("/users/:id", handlers.DeleteUser(conn))
	app.Post("/register", handlers.RegisterUser(conn))

	// Auth Routes for User
	app.Post("/login", handlers.Login(conn))

	//Middleware Routes test
	app.Get("/me", middleware.Protect(), func(req fiber.Ctx) error {
		userID := req.Locals("user_id")
		email := req.Locals("email")
		return req.Status(200).JSON(fiber.Map{"user_id": userID, "email": email})
	})
}
