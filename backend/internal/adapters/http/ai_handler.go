package http

import (
	"encoding/json"
	"net/http"
)

// GenerateAssignmentRequest represents the request payload for AI assignment generation.
type GenerateAssignmentRequest struct {
	Class   string `json:"class"`
	Subject string `json:"subject"`
	Prompt  string `json:"prompt"`
}

// Question represents an individual generated question within an assignment.
type Question struct {
	Question string `json:"question"`
	Type     string `json:"type"`
	Marks    int    `json:"marks"`
}

// GenerateAssignmentResponse represents the response payload for AI assignment generation.
type GenerateAssignmentResponse struct {
	Title        string     `json:"title"`
	Instructions string     `json:"instructions"`
	Questions    []Question `json:"questions"`
	TotalMarks   int        `json:"total_marks"`
}

// GenerateAssignment handles AI assignment generation requests.
func (h *Handler) GenerateAssignment(w http.ResponseWriter, r *http.Request) {
	var req GenerateAssignmentRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Validate required fields
	if req.Class == "" || req.Subject == "" || req.Prompt == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "class, subject, and prompt are required"})
		return
	}

	generated, err := h.aiService.GenerateAssignment(r.Context(), req.Class, req.Subject, req.Prompt)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	questions := make([]Question, 0, len(generated.Questions))
	for _, q := range generated.Questions {
		questions = append(questions, Question{
			Question: q.Question,
			Type:     q.Type,
			Marks:    q.Marks,
		})
	}

	resp := GenerateAssignmentResponse{
		Title:        generated.Title,
		Instructions: generated.Instructions,
		Questions:    questions,
		TotalMarks:   generated.TotalMarks,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
