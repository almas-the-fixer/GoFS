package handlers

import (
	"fmt"
	"gofs/internal/database"
	"gofs/internal/types"
	"gofs/internal/validation"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

// Get One User
func GetUser(conn *pgx.Conn) fiber.Handler {
	return func(c fiber.Ctx) error {
		userID := c.Params("id")

		user, err := database.GetUser(conn, userID)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "An Error Occured"})
		}
		return c.Status(200).JSON(fiber.Map{"status": "OK", "id": user.ID, "user": user})
	}
}

// Get All Users (Admin Only Soon!)
func GetUsers(conn *pgx.Conn) fiber.Handler {
	return func(c fiber.Ctx) error {
		users := database.GetUsers(conn)
		return c.JSON(users)
	}
}

// Delete a User (TODO Admin Only after Middleware!)
func DeleteUser(conn *pgx.Conn) fiber.Handler {
	return func(req fiber.Ctx) error {
		userID := req.Params("id")
		err := database.DeleteUser(conn, userID)
		if err != nil {
			return req.Status(400).JSON(fiber.Map{"Status": "Bad Request"})
		}
		fmt.Println("Deleted Successfullly")
		return nil
	}
}

// Create a User
func CreateUser(conn *pgx.Conn) fiber.Handler {
	return func(req fiber.Ctx) error {
		user := new(types.UserCreateRequest)
		if err := req.Bind().Body(user); err != nil {
			return fiber.ErrBadRequest
		}
		err := validation.UserCreateRequestValidator(*user)
		if err != nil {
			return err
		}
		// Hashing Password Before Storing it
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("An Error Occures while HASHING password: ", err)
			return err
		}
		
		// Need to change this because its mutating the fields and user.Password now contains hashedPassword, that can be confusing so later on needs refactor
		user.Password = string(hashedPass)
		id, err := database.InsertUser(conn, *user)
		if err != nil {
			fmt.Println("An Error Occured: ", err)
			return req.Status(401).JSON(fiber.Map{"status": "Bad request"})
		}
		return req.Status(201).JSON(fiber.Map{"status": "user created successfully", "user_id": id})
	}
}
