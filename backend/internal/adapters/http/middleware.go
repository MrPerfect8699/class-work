package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

// AuthMiddleware validates JWT tokens
func (h *Handler) AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if auth == "" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": "missing authorization header"})
				return
			}

			tokenStr := strings.TrimPrefix(auth, "Bearer ")
			if tokenStr == auth {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": "invalid authorization header format"})
				return
			}

			teacherID, err := h.authService.ValidateToken(tokenStr)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": "invalid token"})
				return
			}

			// Store teacherID in request context
			ctx := context.WithValue(r.Context(), "teacherID", teacherID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Health returns the health status of the server
func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}
