package main

import (
	"log"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "healthy"})
	})

	log.Fatal(app.Listen(":3000"))
}
