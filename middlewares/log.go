package middlewares

import (
	"log"
	"net/http"
)

/**
Log Middleware
*/
func Log(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %v", r.URL)
		handler.ServeHTTP(w, r)
	}
}
