package main

import (
	"fmt"
	"go-user-management-service/src/internal/api/routers"
	"go-user-management-service/src/internal/config"
	"go-user-management-service/src/internal/logger"
	"net/http"
	"runtime"
)

var _, file, line, _ = runtime.Caller(0)
var log = logger.SetupLogger(file, line)

func main() {
	log.Infof("Server starting on %s:%d...", config.AppConfig.HOST, config.AppConfig.PORT)
	apiRouter := routers.GetRouter()

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", config.AppConfig.HOST, config.AppConfig.PORT), apiRouter)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
