package handlers

import (
	"main/api"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/addgame", api.AddGames)

}
