package routers

import (
	"github.com/gorilla/mux"
	"go-user-management-service/src/internal/api/middlewares"
)

func GetRouter() *mux.Router {
	apiRouter := mux.NewRouter()
	apiRouter.PathPrefix("/").Handler(GetHealthCheckRouter())

	apiRouter.Use(middlewares.LoggingMiddleware())
	return apiRouter
}
