package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"time"
)

var AppConfig *Config

type Config struct {
	LOGLEVEL           string
	HOST               string
	PORT               int
	ELASTICSEARCH_HOST string
	ELASTICSEARCH_PORT int
	APP_VERSION        string
	APP_HOST           string
}

func init() {
	AppConfig = &Config{}
	AppConfig.setupConfig()
}

func (c *Config) setupConfig() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("can't load env vars")
	} else {
		err = viper.Unmarshal(&c)
		if err != nil {
			log.Printf("can't load env vars")
		}
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		go c.handleConfigChangeWithDelay()
	})
}

func (c *Config) handleConfigChangeWithDelay() {
	log.Printf("Change in env file detected, reloading config...")
	time.Sleep(2 * time.Second)
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Printf("can't load env vars")
	}
	log.Printf("Config reloaded")
}

func (c *Config) getEnv(key string, defaultValue interface{}) interface{} {
	if !viper.IsSet(key) {
		return defaultValue
	}
	return viper.Get(key)
}
