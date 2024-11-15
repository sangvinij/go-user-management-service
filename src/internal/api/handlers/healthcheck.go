package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-user-management-service/src/internal/api/schemas/healthcheck"
)

// GetHealthCheck проверяет состояние приложения
//
//	@Summary		Health Check
//
//	@Description	Check if the service is healthy
//	@Tags			HealthCheck
//	@Produce		json
//	@Success		200	{object}	healthcheck.ViewModelHealthCheck
//	@Router			/um/healthcheck [get]
func GetHealthCheck(context *fiber.Ctx) error {
	response := healthcheck.ViewModelHealthCheck{
		Status: "healthy",
	}
	return context.Status(fiber.StatusOK).JSON(response)
}
