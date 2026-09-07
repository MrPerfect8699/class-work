package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/yourname/classwork/backend/internal/domain"
)

type mockAIService struct {
	generateFn func(ctx context.Context, className, subject, prompt string) (*domain.GeneratedAssignment, error)
}

func (m *mockAIService) GenerateAssignment(ctx context.Context, className, subject, prompt string) (*domain.GeneratedAssignment, error) {
	if m.generateFn != nil {
		return m.generateFn(ctx, className, subject, prompt)
	}
	return &domain.GeneratedAssignment{
		Title:        "Quadratic Equations - Practice Set",
		Instructions: "Solve all questions and show your working.",
		Questions: []domain.GeneratedQuestion{
			{
				Question: "Solve x² - 5x + 6 = 0",
				Type:     "short_answer",
				Marks:    2,
			},
		},
		TotalMarks: 20,
	}, nil
}

func TestGenerateAssignmentRequestJSON(t *testing.T) {
	reqJSON := `{
		"class": "Grade 10-A",
		"subject": "Mathematics",
		"prompt": "Create 10 questions on quadratic equations"
	}`

	var req GenerateAssignmentRequest
	if err := json.Unmarshal([]byte(reqJSON), &req); err != nil {
		t.Fatalf("failed to unmarshal GenerateAssignmentRequest: %v", err)
	}

	if req.Class != "Grade 10-A" {
		t.Errorf("expected Class 'Grade 10-A', got %q", req.Class)
	}
	if req.Subject != "Mathematics" {
		t.Errorf("expected Subject 'Mathematics', got %q", req.Subject)
	}
	if req.Prompt != "Create 10 questions on quadratic equations" {
		t.Errorf("expected Prompt 'Create 10 questions on quadratic equations', got %q", req.Prompt)
	}
}

func TestGenerateAssignmentResponseJSON(t *testing.T) {
	resp := GenerateAssignmentResponse{
		Title:        "Quadratic Equations - Practice Set",
		Instructions: "Solve all questions and show your working.",
		Questions: []Question{
			{
				Question: "Solve x² - 5x + 6 = 0",
				Type:     "short_answer",
				Marks:    2,
			},
		},
		TotalMarks: 20,
	}

	data, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("failed to marshal GenerateAssignmentResponse: %v", err)
	}

	var decoded GenerateAssignmentResponse
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal GenerateAssignmentResponse: %v", err)
	}

	if decoded.Title != resp.Title {
		t.Errorf("expected Title %q, got %q", resp.Title, decoded.Title)
	}
	if decoded.Instructions != resp.Instructions {
		t.Errorf("expected Instructions %q, got %q", resp.Instructions, decoded.Instructions)
	}
	if len(decoded.Questions) != 1 {
		t.Fatalf("expected 1 question, got %d", len(decoded.Questions))
	}
	if decoded.Questions[0].Question != "Solve x² - 5x + 6 = 0" {
		t.Errorf("expected Question 'Solve x² - 5x + 6 = 0', got %q", decoded.Questions[0].Question)
	}
	if decoded.Questions[0].Type != "short_answer" {
		t.Errorf("expected Type 'short_answer', got %q", decoded.Questions[0].Type)
	}
	if decoded.Questions[0].Marks != 2 {
		t.Errorf("expected Marks 2, got %d", decoded.Questions[0].Marks)
	}
	if decoded.TotalMarks != 20 {
		t.Errorf("expected TotalMarks 20, got %d", decoded.TotalMarks)
	}
}

func TestGenerateAssignmentEndpoint_Success(t *testing.T) {
	handler := &Handler{
		aiService: &mockAIService{},
	}

	requestBody := `{"class": "Grade 10-A", "subject": "Mathematics", "prompt": "Create 10 questions on quadratic equations"}`
	req := httptest.NewRequest(http.MethodPost, "/api/ai/assignments", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.GenerateAssignment(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d. Body: %s", w.Code, w.Body.String())
	}

	var resp GenerateAssignmentResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Title == "" {
		t.Error("expected non-empty title in response")
	}
	if len(resp.Questions) == 0 {
		t.Error("expected questions in response")
	}
	if resp.TotalMarks <= 0 {
		t.Errorf("expected positive total_marks, got %d", resp.TotalMarks)
	}
}

func TestGenerateAssignmentEndpoint_ServiceError(t *testing.T) {
	handler := &Handler{
		aiService: &mockAIService{
			generateFn: func(ctx context.Context, className, subject, prompt string) (*domain.GeneratedAssignment, error) {
				return nil, errors.New("ai generation service unavailable")
			},
		},
	}

	requestBody := `{"class": "Grade 10-A", "subject": "Mathematics", "prompt": "Create 10 questions on quadratic equations"}`
	req := httptest.NewRequest(http.MethodPost, "/api/ai/assignments", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.GenerateAssignment(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500 Internal Server Error, got %d. Body: %s", w.Code, w.Body.String())
	}
}

func TestGenerateAssignmentEndpoint_ValidationErrors(t *testing.T) {
	handler := &Handler{
		aiService: &mockAIService{},
	}

	testCases := []struct {
		name        string
		body        string
		expectedErr string
	}{
		{
			name:        "Invalid JSON",
			body:        `{invalid json`,
			expectedErr: "invalid character",
		},
		{
			name:        "Missing class",
			body:        `{"class": "", "subject": "Math", "prompt": "Test"}`,
			expectedErr: "class, subject, and prompt are required",
		},
		{
			name:        "Missing subject",
			body:        `{"class": "10-A", "subject": "", "prompt": "Test"}`,
			expectedErr: "class, subject, and prompt are required",
		},
		{
			name:        "Missing prompt",
			body:        `{"class": "10-A", "subject": "Math", "prompt": ""}`,
			expectedErr: "class, subject, and prompt are required",
		},
		{
			name:        "Empty object",
			body:        `{}`,
			expectedErr: "class, subject, and prompt are required",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/ai/assignments", bytes.NewBufferString(tc.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			handler.GenerateAssignment(w, req)

			if w.Code != http.StatusBadRequest {
				t.Fatalf("expected status 400 Bad Request, got %d. Body: %s", w.Code, w.Body.String())
			}

			var errResp map[string]string
			if err := json.NewDecoder(w.Body).Decode(&errResp); err != nil {
				t.Fatalf("failed to decode error response: %v", err)
			}

			if errMsg, ok := errResp["error"]; !ok || errMsg == "" {
				t.Errorf("expected error field in response, got %v", errResp)
			}
		})
	}
}

func TestGenerateAssignmentEndpoint_RouterIntegration(t *testing.T) {
	handler := NewHandler(nil, nil, &mockAIService{})
	r := chi.NewRouter()
	handler.RegisterRoutes(r)

	requestBody := `{"class": "Grade 10-A", "subject": "Mathematics", "prompt": "Create 10 questions on quadratic equations"}`
	req := httptest.NewRequest(http.MethodPost, "/api/ai/assignments", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200 OK via router, got %d. Body: %s", w.Code, w.Body.String())
	}
}
