package main

import (
	"fmt"
	"go-user-management-service/src/internal/api/routers"
	"go-user-management-service/src/internal/config"
	"go-user-management-service/src/internal/logger"
	"net/http"
)

func main() {
	cfg, err := config.GetConfig()

	log := logger.SetupLogger(cfg)

	log.Infof("Server starting on %s:%d...", cfg.Host, cfg.Port)
	apiRouter := routers.GetRouter(log)

	err = http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), apiRouter)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
