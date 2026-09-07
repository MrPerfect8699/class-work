package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/yourname/classwork/backend/internal/domain"
)

// CreateHomeworkRequest represents a create homework request
type CreateHomeworkRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ClassName   string `json:"className"`
	Subject     string `json:"subject"`
	TeacherID   int64  `json:"teacherId,omitempty"`
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

	var teacherID int64
	if ctxTeacherID, ok := r.Context().Value("teacherID").(int64); ok && ctxTeacherID > 0 {
		teacherID = ctxTeacherID
	} else if req.TeacherID > 0 {
		teacherID = req.TeacherID
	}

	if teacherID <= 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "teacher_id not found in context or request"})
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
	var teacherID int64
	if ctxTeacherID, ok := r.Context().Value("teacherID").(int64); ok && ctxTeacherID > 0 {
		teacherID = ctxTeacherID
	}

	// Support query param ?teacherId=...
	if queryTeacherID := r.URL.Query().Get("teacherId"); queryTeacherID != "" {
		if parsedID, err := strconv.ParseInt(queryTeacherID, 10, 64); err == nil && parsedID > 0 {
			teacherID = parsedID
		}
	}

	if teacherID <= 0 {
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
	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid homework_id"})
		return
	}

	homework, err := h.homeworkService.GetHomeworkByID(id)
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
	idStr := chi.URLParam(r, "id")
	var req UpdateHomeworkRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid homework_id"})
		return
	}

	existing, err := h.homeworkService.GetHomeworkByID(id)
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
	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid homework_id"})
		return
	}

	err = h.homeworkService.DeleteHomework(id)
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
