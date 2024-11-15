package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-user-management-service/src/internal/api/middlewares"
	"go-user-management-service/src/internal/api/routers"
	"go-user-management-service/src/internal/config"
	"go-user-management-service/src/internal/logger"
	"runtime"
)

var _, file, line, _ = runtime.Caller(0)
var log = logger.SetupLogger(file, line)

func main() {
	app := fiber.New()
	app.Use(middlewares.LoggerMiddleware)

	routers.SetupRouter(app)

	log.Infof("Server starting on %s:%d...", config.AppConfig.HOST, config.AppConfig.PORT)
	err := app.Listen(fmt.Sprintf("%s:%d", config.AppConfig.HOST, config.AppConfig.PORT))
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
