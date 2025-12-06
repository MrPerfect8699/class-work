package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
)

func main() {
    // load env from environment (in docker-compose we'll set them)
    ConnectDB()
    defer DB.Close()

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
