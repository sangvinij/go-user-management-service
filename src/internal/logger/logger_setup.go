package logger

import (
	"github.com/sirupsen/logrus"
	"go-user-management-service/src/internal/integrations/elasticsearch"
	"os"
)

func SetupLogger(file string, line int) *logrus.Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{})

	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.SetLevel(logrus.InfoLevel)
	} else {
		log.SetLevel(level)
	}

	esClient, err := elasticsearch.NewElasticClient()
	if err != nil {
		log.Fatalf("elasticsearch client error: %v", err)
	}

	log.AddHook(GetLogHook(file, line))

	hook, err := NewElasticHook(esClient, "go-logs")
	if err != nil {
		log.Fatalf("elasticsearch hook creation error: %v", err)
	}
	log.AddHook(hook)
	return log
}
