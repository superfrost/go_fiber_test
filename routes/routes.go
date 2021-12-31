package routes

import (
	"fiber_test/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)

	app.Post("/api/login", controllers.Login)

	// app.Get("/env")

	// app.Get("/port")

	// app.Get("/:name")
}

// app.Get("/env", func(c *fiber.Ctx) error {
// 	return c.SendString("Env var is: " + os.Getenv("MES"))
// })

// app.Get("/port", func(c *fiber.Ctx) error {
// 	return c.SendString("PORT is: " + os.Getenv("PORT"))
// })

// app.Get("/:name", func(c *fiber.Ctx) error {
// 	msg := fmt.Sprintf("Hello, %s! ", c.Params("name"))
// 	return c.SendString(msg + fmt.Sprint(time.Now().Unix()))
// })
