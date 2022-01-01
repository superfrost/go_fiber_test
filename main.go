package main

import (
	"fiber_test/database"
	"fiber_test/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Static("/", "./public")

	routes.Setup(app)

	app.Listen(":" + os.Getenv("PORT"))
}
