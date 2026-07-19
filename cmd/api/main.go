package main

import (
	"context"
	"fmt"
	"log"

	"gofs/internal/database"
	"gofs/internal/types"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

	func main() {

		// COMPOSITION ROOT

		//Getting Environment Variables
		app := fiber.New()
		if err := godotenv.Load(); err != nil {
			fmt.Println("AN ERROR OCCURED LOADING THE ENV VARIABLES: ", err)
		}

		// Connection to PostgresDB
		conn , err := database.ConnectDB()
		if err != nil {
			log.Fatal("AN ERROR OCCURED WHEN CONNECTING TO DB: ", err)
		}
		
		defer conn.Close(context.Background())
		
		app.Get("/health", func(c fiber.Ctx) error {
			return c.JSON(fiber.Map{"status": "healthy"})
		})

		app.Get("/", func(c fiber.Ctx) error {
			return c.SendString("Hello, World!")
		})

		// This will be protected as an Admin only Route in Future
		app.Get("/users", func (req fiber.Ctx) error {
			users := database.GetUsers(conn)
			return req.JSON(users)
		})

		app.Post("/users", func(req fiber.Ctx) error {
			user := new(types.UserCreateRequest)
			if err := req.Bind().Body(user); err != nil {
				return fiber.ErrBadRequest
			}
			err := database.CreateUser(conn, *user)
			if err != nil {
				fmt.Println("An Error Occured: ", err)
			}
			return req.Status(200).JSON(fiber.Map{"status": "user inserted successfully"})
		})

		log.Fatal(app.Listen(":3000"))
	}
