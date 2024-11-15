package openapi

import (
	"github.com/gofiber/swagger"
	"go-user-management-service/src/docs"
	"go-user-management-service/src/internal/config"
)

func configureSwaggerInfo() {
	docs.SwaggerInfo.Version = config.AppConfig.APP_VERSION
	docs.SwaggerInfo.Host = config.AppConfig.APP_HOST
}

func SwaggerConfig() swagger.Config {
	configureSwaggerInfo()
	cfg := swagger.Config{
		Title: "User Management",
	}
	return cfg
}
