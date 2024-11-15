package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetHealthCheck(context *fiber.Ctx) error {
	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "healthy",
	})
}
