package routers

import (
	"github.com/gofiber/fiber/v2"
	"go-user-management-service/src/internal/api/middlewares"
)

func SetupRouter(app *fiber.App) {
	baseRouter := app.Group("/um")
	baseRouter.Use(middlewares.LoggerMiddleware)
	GetHealthCheckRouter(baseRouter)
}
