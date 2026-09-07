package ports

import "github.com/yourname/classwork/backend/internal/domain"

// AuthService defines the authentication service interface
type AuthService interface {
	Register(teacher *domain.Teacher) (*domain.Teacher, error)
	Login(email, password string) (string, error)
	ValidateToken(token string) (int64, error)
}
