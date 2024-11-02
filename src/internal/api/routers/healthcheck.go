package routers

import (
	"github.com/gorilla/mux"
	"go-user-management-service/src/internal/api/handlers"
)

func GetHealthCheckRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", handlers.GetHealthCheck).Methods("GET")
	return router
}
