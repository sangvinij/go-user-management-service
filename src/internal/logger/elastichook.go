package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/sirupsen/logrus"
	"time"
)

type LogMessage struct {
	Level              string    `json:"level"`
	Message            string    `json:"message"`
	File               string    `json:"file,omitempty"`
	Line               int       `json:"line,omitempty"`
	Timestamp          time.Time `json:"@timestamp"`
	RequestBody        string    `json:"request_body,omitempty"`
	RequestUrl         string    `json:"request_url,omitempty"`
	RequestMethod      string    `json:"request_method,omitempty"`
	ResponseBody       string    `json:"response_body,omitempty"`
	ResponseStatusCode int       `json:"response_status,omitempty"`
}

type ElasticHook struct {
	client *elasticsearch.Client
	index  string
}

func NewElasticHook(client *elasticsearch.Client, index string) (*ElasticHook, error) {
	res, err := client.Ping()
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, fmt.Errorf("elasticsearch status error: %s", res.String())
	}

	return &ElasticHook{
		client: client,
		index:  index,
	}, err
}

func (hook *ElasticHook) Fire(entry *logrus.Entry) error {
	logMsg := LogMessage{
		Level:     entry.Level.String(),
		Message:   entry.Message,
		Timestamp: entry.Time,
	}

	if file, ok := entry.Data["file"].(string); ok {
		logMsg.File = file
	}
	if line, ok := entry.Data["line"].(int); ok {
		logMsg.Line = line
	}

	if reqURL, ok := entry.Data["request_url"].(string); ok {
		logMsg.RequestUrl = reqURL
	}

	if reqMethod, ok := entry.Data["request_method"].(string); ok {
		logMsg.RequestMethod = reqMethod
	}

	if reqBody, ok := entry.Data["request_body"].(string); ok {
		logMsg.RequestBody = reqBody
	}

	if resStatusCode, ok := entry.Data["response_status"].(int); ok {
		logMsg.ResponseStatusCode = resStatusCode
	}
	if resBody, ok := entry.Data["response_body"].(string); ok {
		logMsg.ResponseBody = resBody
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(logMsg); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	res, err := hook.client.Index(hook.index, &buf)
	if err != nil {
		return fmt.Errorf("error while sending logs to Elasticsearch: %w", err)
	}

	if res.IsError() {
		return fmt.Errorf("indexing Error in Elasticsearch: %s", res.String())
	}

	return nil
}

func (hook *ElasticHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
