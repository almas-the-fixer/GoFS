	package main

	import (
		"context"
		"fmt"
		"log"

		"gofs/internal/database"

		"github.com/gofiber/fiber/v3"
		"github.com/joho/godotenv"
	)

	func main() {
		app := fiber.New()
		if err := godotenv.Load(); err != nil {
			fmt.Println("AN ERROR OCCURED LOADING THE ENV VARIABLES: ", err)
		}

		conn , err := database.ConnectDB()
		if err != nil {
			log.Fatal("AN ERROR OCCURED WHEN CONNECTING TO DB: ", err)
		}
		
		defer conn.Close(context.Background())



		app.Get("/", func(c fiber.Ctx) error {
			return c.SendString("Hello, World!")
		})

		app.Get("/users", func (c fiber.Ctx) error {
			users := database.GetUsers(conn)
			return c.JSON(users)
		})

		app.Get("/health", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{"status": "healthy"})
		})

		log.Fatal(app.Listen(":3000"))
	}
