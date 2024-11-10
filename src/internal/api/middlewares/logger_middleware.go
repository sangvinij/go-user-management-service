package middlewares

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"go-user-management-service/src/internal/logger"
	"io"
	"net/http"
	"runtime"
)

var _, file, line, _ = runtime.Caller(0)
var log = logger.SetupLogger(file, line)

type CustomResponseWriter struct {
	http.ResponseWriter
	status int
	body   *bytes.Buffer
}

func (rw *CustomResponseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

func (rw *CustomResponseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b)
	n, err := rw.ResponseWriter.Write(b)
	return n, err
}

func LoggingMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rw := &CustomResponseWriter{ResponseWriter: w, status: http.StatusOK, body: &bytes.Buffer{}}

			bodyBytes, _ := io.ReadAll(r.Body)
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			log.WithFields(logrus.Fields{
				"request_url":    r.URL.Path,
				"request_method": r.Method,
				"request_body":   string(bodyBytes),
			}).Info("Receive request from client")

			next.ServeHTTP(rw, r)

			log.WithFields(logrus.Fields{
				"response_status": rw.status,
				"response_body":   rw.body.String(),
			}).Info("Send response to client")
		})
	}
}
