package application

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/yourname/classwork/backend/config"
	"github.com/yourname/classwork/backend/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// AuthServiceImpl implements the AuthService interface
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

// Register registers a new teacher
func (s *AuthServiceImpl) Register(name, email, password string) (primitive.ObjectID, error) {
	// Check if email already exists
	_, err := s.teacherRepo.FindByEmail(email)
	if err == nil {
		return primitive.NilObjectID, fmt.Errorf("email already taken")
	}

	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create teacher
	teacher := &domain.Teacher{
		ID:        primitive.NewObjectID(),
		Name:      name,
		Email:     email,
		Password:  string(hashed),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.teacherRepo.Save(teacher)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("failed to save teacher: %w", err)
	}

	return teacher.ID, nil
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
		"teacher_id": teacher.ID.Hex(),
		"email":      teacher.Email,
		"exp":        time.Now().Add(s.config.JWT.ExpirationTime).Unix(),
		"iat":        time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.config.JWT.Secret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

// ValidateToken validates a JWT token and returns the teacher ID
func (s *AuthServiceImpl) ValidateToken(tokenString string) (primitive.ObjectID, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.JWT.Secret), nil
	})

	if err != nil || !token.Valid {
		return primitive.NilObjectID, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("invalid token claims")
	}

	teacherIDHex, ok := claims["teacher_id"].(string)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("invalid teacher_id in token")
	}

	teacherID, err := primitive.ObjectIDFromHex(teacherIDHex)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("invalid teacher_id format")
	}

	return teacherID, nil
}
