package main

import (
	"sederhana/service"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()
	app.Get("/", helloWorld)
	app.Post("/", service.Create)
	app.Put("/:id", service.Update)
	app.Delete("/:id", service.Delete)
	app.Get("/", service.Get)
	app.Listen(":3000")
}

func helloWorld(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"message": "api working"})
}
