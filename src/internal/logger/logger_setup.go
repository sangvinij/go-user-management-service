package logger

import (
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
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

	elasticURL := "http://localhost:9200"
	client, err := elastic.NewClient(elastic.SetURL(elasticURL), elastic.SetSniff(false))
	if err != nil {
		log.Fatalf("Ошибка при подключении к Elasticsearch: %v", err)
	}

	hook, err := elogrus.NewAsyncElasticHook(client, "localhost", level, "go-logs")
	if err != nil {
		log.Fatalf("Ошибка при создании Elasticsearch хука: %v", err)
	}
	log.AddHook(hook)
	log.AddHook(GetLogHook(file, line))

	return log
}
