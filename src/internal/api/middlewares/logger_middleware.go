package middlewares

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func LoggingMiddleware(log *logrus.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.WithFields(logrus.Fields{
				"request_url":  r.URL.Path,
				"method":       r.Method,
				"request_body": r.Form.Encode(),
			}).Info("Receive request")
			next.ServeHTTP(w, r)
		})
	}
}
