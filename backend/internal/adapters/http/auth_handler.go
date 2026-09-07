package http

import (
	"encoding/json"
	"net/http"

	"github.com/yourname/classwork/backend/internal/domain"
)

// RegisterRequest represents a registration request
type RegisterRequest struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	TeacherID       string `json:"teacherId,omitempty"`
	Mobile          string `json:"mobile,omitempty"`
	Department      string `json:"department,omitempty"`
	Designation     string `json:"designation,omitempty"`
	Qualification   string `json:"qualification,omitempty"`
	ExperienceYears int    `json:"experienceYears,omitempty"`
	AvatarURL       string `json:"avatarUrl,omitempty"`
}

// LoginRequest represents a login request
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthResponse represents the authentication response
type AuthResponse struct {
	ID          int64  `json:"id,omitempty"`
	TeacherID   string `json:"teacherId,omitempty"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	Mobile      string `json:"mobile,omitempty"`
	Department  string `json:"department,omitempty"`
	Designation string `json:"designation,omitempty"`
	Status      string `json:"status,omitempty"`
	Token       string `json:"token,omitempty"`
	Error       string `json:"error,omitempty"`
}

// Register handles teacher registration
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Validate required fields
	if req.Name == "" || req.Email == "" || req.Password == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "name, email, and password are required"})
		return
	}

	teacher := &domain.Teacher{
		TeacherID:       req.TeacherID,
		Name:            req.Name,
		Email:           req.Email,
		Mobile:          req.Mobile,
		Password:        req.Password,
		Department:      req.Department,
		Designation:     req.Designation,
		Qualification:   req.Qualification,
		ExperienceYears: req.ExperienceYears,
		AvatarURL:       req.AvatarURL,
	}

	createdTeacher, err := h.authService.Register(teacher)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(AuthResponse{
		ID:          createdTeacher.ID,
		TeacherID:   createdTeacher.TeacherID,
		Name:        createdTeacher.Name,
		Email:       createdTeacher.Email,
		Mobile:      createdTeacher.Mobile,
		Department:  createdTeacher.Department,
		Designation: createdTeacher.Designation,
		Status:      createdTeacher.Status,
	})
}

// Login handles teacher login
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Validate required fields
	if req.Email == "" || req.Password == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "email and password are required"})
		return
	}

	token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(AuthResponse{Token: token})
}
