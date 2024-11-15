package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go-user-management-service/src/internal/logger"
	"runtime"
)

var _, file, line, _ = runtime.Caller(0)
var log = logger.SetupLogger(file, line)

func LoggerMiddleware(c *fiber.Ctx) error {

	log.WithFields(logrus.Fields{
		"request_url":    string(c.Request().RequestURI()),
		"request_method": c.Method(),
		"request_body":   string(c.Request().Body()),
	}).Info("Receive request from client")

	err := c.Next()

	log.WithFields(logrus.Fields{
		"response_status": c.Response().StatusCode(),
		"response_body":   string(c.Response().Body()),
	}).Info("Send response to client")

	return err
}
