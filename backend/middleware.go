package main

import (
    "net/http"
    "os"
    "strings"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        auth := c.GetHeader("Authorization")
        if auth == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing auth"})
            return
        }
        tokenStr := strings.TrimPrefix(auth, "Bearer ")
        secret := os.Getenv("JWT_SECRET")
        token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
            return []byte(secret), nil
        })
        if err != nil || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            return
        }
        c.Next()
    }
}
