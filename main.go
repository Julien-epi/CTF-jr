package main

import (
	"CTF-JR/entrypoint" 
	"CTF-JR/routes"
	"fmt"  
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	log.Info("Main function started")

	port, err := entrypoint.CurrentPort()
	if err != nil {
		log.Fatal(err) 
	}

	fmt.Println("Dynamic port found:", port)

	app := fiber.New(fiber.Config{})
	routes.SetupRoutes(app)
	app.Listen(":81")
	log.Info("Starting server...")
}
