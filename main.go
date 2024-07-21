package main

import (		
	"fiber-rest-api/database"
	"fiber-rest-api/database/migrations"
	"fiber-rest-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Inisialisasi Database
	database.DatabaseInit()

	// Inisialisasi Migration
	migrations.Migration()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Hello",
		})
	})

	// Inisialisasi Rute
	routes.RouteInit(app)

	app.Listen(":8080")
}
