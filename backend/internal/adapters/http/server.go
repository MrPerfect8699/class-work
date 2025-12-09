package http

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/yourname/classwork/backend/config"
)

// Server represents the HTTP server
type Server struct {
	router chi.Router
	config *config.ServerConfig
	http   *http.Server
}

// NewServer creates a new HTTP server with Chi router and CORS
func NewServer(handler *Handler, cfg *config.ServerConfig) *Server {
	router := chi.NewRouter()

	// Add Chi middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)

	// Add CORS middleware with localhost:* allowed
	router.Use(cors.Handler(DefaultCORSConfig()))

	// Register routes
	handler.RegisterRoutes(router)

	return &Server{
		router: router,
		config: cfg,
	}
}

// Start starts the HTTP server
func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)

	s.http = &http.Server{
		Addr:         addr,
		Handler:      s.router,
		ReadTimeout:  s.config.ReadTimeout,
		WriteTimeout: s.config.WriteTimeout,
		IdleTimeout:  s.config.IdleTimeout,
	}

	log.Printf("Starting server on %s", addr)

	if err := s.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("server error: %w", err)
	}

	return nil
}

// Shutdown gracefully shuts down the HTTP server
func (s *Server) Shutdown(ctx context.Context) error {
	if s.http == nil {
		return nil
	}

	return s.http.Shutdown(ctx)
}

// GetRouter returns the Chi router
func (s *Server) GetRouter() chi.Router {
	return s.router
}
