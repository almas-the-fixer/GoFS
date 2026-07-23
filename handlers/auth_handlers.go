package handlers

import (
	"fmt"
	"gofs/internal/database"
	"gofs/internal/types"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)


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