package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		path := r.URL.Path
		method := r.Method
		fmt.Printf("[start]time:%s, method:%s, path:%s\n", start.Format(time.RFC3339Nano), method, path)
		next.ServeHTTP(w, r)
		end := time.Now()
		elapsedTime := end.Sub(start).Milliseconds()
		fmt.Printf("[END]time:%s, method:%s, path:%s, elapsed_time:%dms\n", end.Format(time.RFC3339Nano), method, path, elapsedTime)
	}
}
