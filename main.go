package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world " + time.Now().String())
	})

	app.Get("/env", func(c *fiber.Ctx) error {
		return c.SendString("Env var is: " + os.Getenv("MES"))
	})

	app.Get("/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, %s! ", c.Params("name"))
		return c.SendString(msg + fmt.Sprint(time.Now().Unix()))
	})

	app.Listen(":" + os.Getenv("PORT"))
}
