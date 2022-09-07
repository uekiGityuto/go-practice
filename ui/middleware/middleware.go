package middleware

import (
	"github.com/uekiGityuto/go-practice/lib/log"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := log.Logger
		//defer logger.Sync()

		start := time.Now()
		path := r.URL.Path
		method := r.Method
		logger.Info("[START]",
			zap.String("method", method),
			zap.String("path", path),
		)

		next.ServeHTTP(w, r)

		end := time.Now()
		elapsedTime := end.Sub(start)
		logger.Info("[END]",
			zap.String("method", method),
			zap.String("method", method),
			zap.String("path", path),
			zap.Duration("elapsed_time(ms)", elapsedTime),
		)
	}
}
