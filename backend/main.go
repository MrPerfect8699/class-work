package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yourname/classwork/backend/config"
	"github.com/yourname/classwork/backend/internal/adapters/http"
	"github.com/yourname/classwork/backend/internal/adapters/persistence"
	"github.com/yourname/classwork/backend/internal/application"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	cfg.LogConfig()

	// Connect to PostgreSQL
	postgresDB, err := persistence.NewPostgresDB(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	log.Println("Connected to PostgreSQL and initialized schema successfully")

	// Initialize repositories
	teacherRepo := persistence.NewTeacherRepository(postgresDB.GetDB())
	homeworkRepo := persistence.NewHomeworkRepository(postgresDB.GetDB())

	// Initialize services
	authService := application.NewAuthService(teacherRepo, cfg)
	homeworkService := application.NewHomeworkService(homeworkRepo)

	// Initialize HTTP handler
	handler := http.NewHandler(authService, homeworkService)

	// Create and start server with Chi router and CORS
	server := http.NewServer(handler, &cfg.Server)

	// Start server in a goroutine
	go func() {
		if err := server.Start(); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for shutdown signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Error shutting down server: %v", err)
	}

	if err := postgresDB.Close(ctx); err != nil {
		log.Printf("Error closing PostgreSQL connection: %v", err)
	}

	log.Println("Server stopped")
}
