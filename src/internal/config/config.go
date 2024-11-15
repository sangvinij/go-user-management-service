package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
)

var AppConfig *config

type config struct {
	LOGLEVEL           string
	HOST               string
	PORT               int
	ELASTICSEARCH_HOST string
	ELASTICSEARCH_PORT int
	APP_VERSION        string
	APP_HOST           string
}

func init() {
	AppConfig = &config{}
	AppConfig.LoadConfig()
}

func (c *config) LoadConfig() {
	godotenv.Load()

	c.LOGLEVEL = c.getEnv("LOG_LEVEL", "info")
	c.HOST = c.getEnv("HOST", "0.0.0.0")
	c.PORT = c.getEnvAsInt("PORT", 8000)
	c.ELASTICSEARCH_HOST = c.getEnv("ELASTICSEARCH_HOST", "http://localhost")
	c.ELASTICSEARCH_PORT = c.getEnvAsInt("ELASTICSEARCH_PORT", 9200)
	c.APP_VERSION = c.getEnv("APP_VERSION", "1.0.0")
	c.APP_HOST = c.getEnv("APP_HOST", "localhost:8000")
}

func (c *config) getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func (c *config) getEnvAsInt(name string, defaultVal int) int {
	valueStr := c.getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func (c *config) getEnvAsBool(name string, defaultVal bool) bool {
	valStr := c.getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return defaultVal
}

func (c *config) getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := c.getEnv(name, "")
	if valStr == "" {
		return defaultVal
	}
	return strings.Split(valStr, sep)
}
