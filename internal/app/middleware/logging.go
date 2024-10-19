package middleware

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

var Logger *zap.SugaredLogger

type Log interface {
	Info(args ...interface{})
}

func InitLogger() {
	logger, _ := zap.NewProduction()
	Logger = logger.Sugar()
}

type responseData struct {
	status int
	size   int
}

type responseWriterWithLogging struct {
	http.ResponseWriter
	responseData *responseData
}

func (r responseWriterWithLogging) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size
	return size, err
}
func (r responseWriterWithLogging) WriteHeader(status int) {
	r.responseData.status = status
	r.ResponseWriter.WriteHeader(status)
}

func LoggingMiddleware(logger *zap.SugaredLogger) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			responseData := &responseData{}
			rwl := responseWriterWithLogging{
				ResponseWriter: w,
				responseData:   responseData,
			}
			start := time.Now()

			h.ServeHTTP(rwl, r)

			duration := time.Since(start)

			logger.Infoln(
				"URL", r.RequestURI,
				"Method:", r.Method,
				"Request Duration:", duration,
				"Status:", responseData.status,
				"Data Size:", responseData.size)
		})
	}
}
