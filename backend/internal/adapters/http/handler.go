package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/yourname/classwork/backend/internal/ports"
)

// Handler holds HTTP request handlers
type Handler struct {
	authService     ports.AuthService
	homeworkService ports.HomeworkService
	aiService       ports.AIService
}

// NewHandler creates a new HTTP handler
func NewHandler(authService ports.AuthService, homeworkService ports.HomeworkService, aiService ports.AIService) *Handler {
	return &Handler{
		authService:     authService,
		homeworkService: homeworkService,
		aiService:       aiService,
	}
}

// RegisterRoutes registers all HTTP routes with Chi router
func (h *Handler) RegisterRoutes(r chi.Router) {
	// Public routes
	r.Post("/api/register", h.Register)
	r.Post("/api/login", h.Login)
	r.Post("/api/ai/assignments", h.GenerateAssignment)

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(h.AuthMiddleware())
		{
			r.Post("/api/homework", h.CreateHomework)
			r.Get("/api/homeworks", h.ListHomeworks)
			r.Get("/api/homework/{id}", h.GetHomework)
			r.Put("/api/homework/{id}", h.UpdateHomework)
			r.Delete("/api/homework/{id}", h.DeleteHomework)
		}
	})

	// Health check
	r.Get("/health", h.Health)
}
