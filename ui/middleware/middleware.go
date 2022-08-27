package middleware

import (
	"fmt"
	"net/http"
)

func Logger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[START] middleware")
		next.ServeHTTP(w, r)
		fmt.Println("[END] middleware")
	}
}
