package routers

import (
	"github.com/gofiber/fiber/v2"
	"go-user-management-service/src/internal/api/handlers"
)

func GetHealthCheckRouter(router fiber.Router) {
	router.Get("/healthcheck", handlers.GetHealthCheck)
}
