package elasticsearch

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"go-user-management-service/src/internal/config"
)

func NewElasticClient() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			fmt.Sprintf("%s:%d", config.AppConfig.ELASTICSEARCH_HOST, config.AppConfig.ELASTICSEARCH_PORT),
		},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return client, nil
}
