package application

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/yourname/classwork/backend/config"
	"github.com/yourname/classwork/backend/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

// AuthServiceImpl implements the ports.AuthService interface
type AuthServiceImpl struct {
	teacherRepo domain.TeacherRepository
	config      *config.Config
}

// NewAuthService creates a new instance of AuthServiceImpl
func NewAuthService(teacherRepo domain.TeacherRepository, cfg *config.Config) *AuthServiceImpl {
	return &AuthServiceImpl{
		teacherRepo: teacherRepo,
		config:      cfg,
	}
}

// Register registers a new teacher with full registration info
func (s *AuthServiceImpl) Register(teacher *domain.Teacher) (*domain.Teacher, error) {
	if teacher.Email == "" || teacher.Name == "" || teacher.Password == "" {
		return nil, fmt.Errorf("name, email, and password are required")
	}

	// Check if email already exists
	_, err := s.teacherRepo.FindByEmail(teacher.Email)
	if err == nil {
		return nil, fmt.Errorf("email already taken")
	}

	// Auto-generate teacher_id code if not provided
	if teacher.TeacherID == "" {
		teacher.TeacherID = fmt.Sprintf("TCH-%d", time.Now().Unix())
	} else {
		// Check if provided teacher_id code already exists
		_, err := s.teacherRepo.FindByTeacherID(teacher.TeacherID)
		if err == nil {
			return nil, fmt.Errorf("teacher_id '%s' is already registered", teacher.TeacherID)
		}
	}

	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(teacher.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	teacher.Password = string(hashed)

	// Set defaults
	if teacher.Status == "" {
		teacher.Status = "ACTIVE"
	}
	if teacher.Designation == "" {
		teacher.Designation = "Teacher"
	}
	teacher.CreatedAt = time.Now()
	teacher.UpdatedAt = time.Now()

	err = s.teacherRepo.Save(teacher)
	if err != nil {
		return nil, fmt.Errorf("failed to save teacher: %w", err)
	}

	return teacher, nil
}

// Login authenticates a teacher and returns a JWT token
func (s *AuthServiceImpl) Login(email, password string) (string, error) {
	// Find teacher by email
	teacher, err := s.teacherRepo.FindByEmail(email)
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(teacher.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"teacher_id":   teacher.ID,
		"teacher_code": teacher.TeacherID,
		"email":        teacher.Email,
		"exp":          time.Now().Add(s.config.JWT.ExpirationTime).Unix(),
		"iat":          time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.config.JWT.Secret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

// ValidateToken validates a JWT token and returns the teacher ID
func (s *AuthServiceImpl) ValidateToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.JWT.Secret), nil
	})

	if err != nil || !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid token claims")
	}

	val, exists := claims["teacher_id"]
	if !exists {
		return 0, fmt.Errorf("invalid teacher_id in token")
	}

	switch v := val.(type) {
	case float64:
		return int64(v), nil
	case int64:
		return v, nil
	case int:
		return int64(v), nil
	case string:
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid teacher_id format")
		}
		return id, nil
	default:
		return 0, fmt.Errorf("invalid teacher_id type in token")
	}
}
