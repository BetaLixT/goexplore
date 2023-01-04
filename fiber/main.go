package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

func main() {
	app := fiber.New()

	app.Use("/", func(c *fiber.Ctx) error {
		fmt.Printf("request %s\n", c.BaseURL())
		c.Next()
		return nil
	})
	app.Use("/", )

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		c.Params("", )
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":8080")
}
