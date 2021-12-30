package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world " + time.Now().String())
	})

	app.Get("/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, %s! ", c.Params("name"))
		return c.SendString(msg + fmt.Sprint(time.Now().Unix()))
	})

	app.Listen(":" + os.Getenv("PORT"))

}
