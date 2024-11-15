package routers

import "github.com/gofiber/fiber/v2"

func SetupRouter(app *fiber.App) {
	baseRouter := app.Group("/um")
	GetHealthCheckRouter(baseRouter)
}
