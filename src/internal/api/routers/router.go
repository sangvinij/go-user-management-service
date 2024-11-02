package routers

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"go-user-management-service/src/internal/api/middlewares"
)

func GetRouter(log *logrus.Logger) *mux.Router {
	apiRouter := mux.NewRouter()
	apiRouter.PathPrefix("/").Handler(GetHealthCheckRouter())

	apiRouter.Use(middlewares.LoggingMiddleware(log))
	return apiRouter
}
