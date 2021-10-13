package main

import (
	"log"

	"github.com/JackWinterburn/go-auth/db"
	"github.com/JackWinterburn/go-auth/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Connect()

	app := fiber.New()

	routes.Setup(app)

	log.Fatal(app.Listen(":8080"))
}
