package middleware

import (
	"net/http"
	"os"
)

var (
	corsOrigin string
)

func init() {
	corsOrigin = os.Getenv("CORS_ORIGIN")
}

// CORS adds necessary headers for CORS
func CORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Origin", corsOrigin)

		h.ServeHTTP(w, r)
		return
	})
}
