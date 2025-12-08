package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load env from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	log.Println("Starting server...")
	ConnectDB() // Ensure database connection is closed on exit
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := DB.Disconnect(ctx); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := gin.Default()

	r.POST("/api/register", Register)
	r.POST("/api/login", Login)

	protected := r.Group("/api")
	protected.Use(AuthMiddleware())
	{
		protected.POST("/homework", CreateHomework)
		protected.GET("/homeworks", ListHomeworks)
	}

	log.Printf("server running on :%s", port)
	r.Run(":" + port)
}
