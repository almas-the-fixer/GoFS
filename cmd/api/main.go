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
	conn, err := database.ConnectDB()
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

	app.Get("/user/:id", func(c fiber.Ctx) error {
		userID := c.Params("id")
		user, err := database.GetUser(conn, userID)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "An Error Occured"})
		}
		return c.Status(200).JSON(fiber.Map{"status": "OK", "id": user.ID, "user": user})
	})

	app.Get("/users", func(req fiber.Ctx) error {
		users := database.GetUsers(conn)
		return req.JSON(users)
	})

	app.Delete("/users/:id", func(req fiber.Ctx) error {
		userID := req.Params("id")
		err := database.DeleteUser(conn, userID)
		if err != nil {
			return req.Status(400).JSON(fiber.Map{"Status": "Bad Request"})
		}
		fmt.Println("Deleted Successfullly")
		return nil
	})

	app.Post("/users", func(req fiber.Ctx) error {
		user := new(types.UserCreateRequest)
		if err := req.Bind().Body(user); err != nil {
			return fiber.ErrBadRequest
		}
		id, err := database.CreateUser(conn, *user)
		if err != nil {
			fmt.Println("An Error Occured: ", err)
			return req.Status(401).JSON(fiber.Map{"status": "Bad request"})
		}
		return req.Status(201).JSON(fiber.Map{"status": "user created successfully", "user_id": id})
	})
	log.Fatal(app.Listen(":3000"))
}
