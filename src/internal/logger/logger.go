package logger

import (
	"github.com/sirupsen/logrus"
	"go-user-management-service/src/internal/config"
	"os"
)

func SetupLogger(cfg *config.Config) *logrus.Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(cfg.LogLevel)
	return log
}
