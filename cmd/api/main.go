package main

import (
	"context"
	"fmt"
	"log"

	"gofs/internal/database"
	"gofs/internal/routes"

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
	// Registering routes
	routes.RegisterRoutes(app, conn)

	log.Fatal(app.Listen(":3000"))
}
