package main 

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app:= fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Connection")
		return c.SendString("hello world fiber")
	})
	app.Get("/api/:type", func(c *fiber.Ctx) error {
		fmt.Println(c.Params("type"))
		return c.SendString("hello world fiber")
	})

	app.Listen(":3000")
}