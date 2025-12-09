package http

import (
	"github.com/go-chi/cors"
)

// CORSConfig defines CORS configuration
type CORSConfig struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	ExposedHeaders   []string
	MaxAge           int
	AllowCredentials bool
}

// DefaultCORSConfig returns a default CORS configuration allowing localhost on any port
func DefaultCORSConfig() cors.Options {
	return cors.Options{
		AllowedOrigins: []string{
			"http://localhost:*",
			"http://127.0.0.1:*",
		},
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"PATCH",
			"OPTIONS",
		},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
			"X-Requested-With",
		},
		ExposedHeaders: []string{
			"Content-Length",
			"X-Request-ID",
		},
		MaxAge:           300,
		AllowCredentials: true,
	}
}

// ProductionCORSConfig returns a CORS configuration for production
func ProductionCORSConfig(allowedOrigins []string) cors.Options {
	return cors.Options{
		AllowedOrigins: allowedOrigins,
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"PATCH",
			"OPTIONS",
		},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
			"X-Requested-With",
		},
		ExposedHeaders: []string{
			"Content-Length",
			"X-Request-ID",
		},
		MaxAge:           3600,
		AllowCredentials: true,
	}
}
