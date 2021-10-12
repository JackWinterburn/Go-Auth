package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	_, err := gorm.Open(postgres.Open(os.Getenv("dsn")), &gorm.Config{})
	if err != nil {
		panic("Database connection failed!")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world")
	})

	log.Fatal(app.Listen(":8080"))
}
