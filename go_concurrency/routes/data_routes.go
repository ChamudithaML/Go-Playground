package routes

import (
	"go_currency/controllers"

	"go_currency/middlewares"

	"github.com/gofiber/fiber/v2"
)

func DataRoute(app *fiber.App) {

	app.Post("/data", middlewares.BufferMiddleware, controllers.AddTemperature)
}
