package main

import (
	"log"
	"main/dbconfig"

	"main/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	dbconfig.SetupDb()

	handlers.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
