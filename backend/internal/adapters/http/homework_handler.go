package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yourname/classwork/backend/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateHomeworkRequest represents a create homework request
type CreateHomeworkRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ClassName   string `json:"className"`
	Subject     string `json:"subject"`
	Attachments string `json:"attachments"`
}

// UpdateHomeworkRequest represents an update homework request
type UpdateHomeworkRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ClassName   string `json:"className"`
	Subject     string `json:"subject"`
	Attachments string `json:"attachments"`
}

// CreateHomework handles homework creation
func (h *Handler) CreateHomework(w http.ResponseWriter, r *http.Request) {
	var req CreateHomeworkRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Validate required fields
	if req.Title == "" || req.ClassName == "" || req.Subject == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "title, className, and subject are required"})
		return
	}

	teacherID, ok := r.Context().Value("teacherID").(primitive.ObjectID)
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "teacher_id not found in context"})
		return
	}

	homework := &domain.Homework{
		Title:       req.Title,
		Description: req.Description,
		ClassName:   req.ClassName,
		Subject:     req.Subject,
		TeacherID:   teacherID,
		Attachments: req.Attachments,
	}

	err := h.homeworkService.CreateHomework(homework)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(homework)
}

// ListHomeworks handles listing homework assignments
func (h *Handler) ListHomeworks(w http.ResponseWriter, r *http.Request) {
	teacherID, ok := r.Context().Value("teacherID").(primitive.ObjectID)
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "teacher_id not found in context"})
		return
	}

	homeworks, err := h.homeworkService.GetHomeworksByTeacher(teacherID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(homeworks)
}

// GetHomework handles getting a specific homework
func (h *Handler) GetHomework(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid homework_id"})
		return
	}

	homework, err := h.homeworkService.GetHomeworkByID(objID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(homework)
}

// UpdateHomework handles updating a homework
func (h *Handler) UpdateHomework(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var req UpdateHomeworkRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid homework_id"})
		return
	}

	existing, err := h.homeworkService.GetHomeworkByID(objID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "homework not found"})
		return
	}

	// Update only provided fields
	if req.Title != "" {
		existing.Title = req.Title
	}
	if req.Description != "" {
		existing.Description = req.Description
	}
	if req.ClassName != "" {
		existing.ClassName = req.ClassName
	}
	if req.Subject != "" {
		existing.Subject = req.Subject
	}
	if req.Attachments != "" {
		existing.Attachments = req.Attachments
	}

	err = h.homeworkService.UpdateHomework(existing)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existing)
}

// DeleteHomework handles deleting a homework
func (h *Handler) DeleteHomework(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid homework_id"})
		return
	}

	err = h.homeworkService.DeleteHomework(objID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "homework deleted successfully"})
}
