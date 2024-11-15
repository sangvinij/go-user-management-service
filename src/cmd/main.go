package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "go-user-management-service/src/docs"
	"go-user-management-service/src/internal/api/openapi"
	"go-user-management-service/src/internal/api/routers"
	"go-user-management-service/src/internal/config"
	"go-user-management-service/src/internal/logger"
	"runtime"
)

var _, file, line, _ = runtime.Caller(0)
var log = logger.SetupLogger(file, line)

// @title			User Management
// @description	go user management service
func main() {
	app := fiber.New()

	app.Get("/docs/*", swagger.New(openapi.SwaggerConfig()))
	routers.SetupRouter(app)

	log.Infof("Server starting on %s:%d...", config.AppConfig.HOST, config.AppConfig.PORT)
	err := app.Listen(fmt.Sprintf("%s:%d", config.AppConfig.HOST, config.AppConfig.PORT))
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
