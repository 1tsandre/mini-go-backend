package middleware

import (
	"net/http"
	"time"

	"github.com/1tsandre/mini-go-backend/pkg/logger"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		logger.Infof("%s %s - %v", r.Method, r.URL.Path, duration)
	})
}
