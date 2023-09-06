package middleware

import (
	"net/http"
	"os"

	"github.com/rs/cors"
)

// Cors ...
func Cors(loggedRoutes http.Handler) http.Handler {
	allowedHeaders := []string{"Authorization", "Content-Type"}
	corsDebug := os.Getenv("CORS_DEBUG")
	if corsDebug == "true" {
		return cors.New(cors.Options{Debug: true, AllowedHeaders: allowedHeaders}).Handler(loggedRoutes)
	}
	return cors.New(cors.Options{AllowedHeaders: allowedHeaders}).Handler(loggedRoutes)
}
