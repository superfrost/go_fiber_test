package main

import (
	"fiber_test/database"
	"fiber_test/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	database.Connect()

	app := fiber.New()

	app.Static("/", "./public")

	routes.Setup(app)

	app.Listen(":" + os.Getenv("PORT"))
}
