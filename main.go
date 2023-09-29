package main

import (
	"CTF-JR/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	log.Info("Main function started")
	app := fiber.New(fiber.Config{})
	routes.SetupRoutes(app)
	app.Listen(":81")
	log.Info("Starting server...")
}