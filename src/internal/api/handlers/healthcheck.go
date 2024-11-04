package handlers

import (
	"go-user-management-service/src/internal/logger"
	"net/http"
	"runtime"
)

var _, file, line, _ = runtime.Caller(0)
var log = logger.SetupLogger(file, line)

func GetHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"status": "ok"}`))
}
