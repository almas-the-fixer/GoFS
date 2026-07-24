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

// Create a User
func RegisterUser(conn *pgx.Conn) fiber.Handler {
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

func Login(conn *pgx.Conn) fiber.Handler{
	return func(req fiber.Ctx) error{
		// Make a type struct login req
		loginReq := new(types.LoginRequest)

		// Parse incoming json req
		if err := req.Bind().Body(loginReq); err != nil {
			return fiber.ErrBadRequest
		}

		// User Exists ?
		dbPassHash, err := database.FindUserByEmail(conn, loginReq.Email)
		if err != nil {
			fmt.Println("Error Occured When Finding user from DB", err)
			return err
		}

		// Compare the passwordHash with dbPassHash
		err = bcrypt.CompareHashAndPassword([]byte(dbPassHash), []byte(loginReq.Password))
		if err != nil{
			fmt.Println("Passsword Doesnt Match with hash!", err)
			return err
		}
		return req.Status(200).JSON(fiber.Map{"status": "login successfull!"})
	}
}