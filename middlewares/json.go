package middlewares

import (
	"log"
	"net/http"
)

/**
Json Middleware
*/
func Json(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %v", r.URL)
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	}
}
