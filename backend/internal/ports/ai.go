package ports

import (
	"context"

	"github.com/yourname/classwork/backend/internal/domain"
)

// AIService defines the AI service port interface for generating assignments.
type AIService interface {
	GenerateAssignment(ctx context.Context, className, subject, prompt string) (*domain.GeneratedAssignment, error)
}
