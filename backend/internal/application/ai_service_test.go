package application

import (
	"context"
	"errors"
	"strings"
	"testing"
)

type mockLLMClient struct {
	generateFn func(ctx context.Context, systemPrompt, userPrompt string) (string, error)
}

func (m *mockLLMClient) Generate(ctx context.Context, systemPrompt, userPrompt string) (string, error) {
	if m.generateFn != nil {
		return m.generateFn(ctx, systemPrompt, userPrompt)
	}
	return `{
		"title": "Quadratic Equations - Practice Set",
		"instructions": "Solve all questions and show your working.",
		"questions": [
			{
				"question": "Solve x² - 5x + 6 = 0",
				"type": "short_answer",
				"marks": 2
			}
		],
		"total_marks": 2
	}`, nil
}

func TestAIServiceImpl_GenerateAssignment_Success(t *testing.T) {
	mockClient := &mockLLMClient{}
	service := NewAIService(mockClient)

	resp, err := service.GenerateAssignment(context.Background(), "Grade 10-A", "Mathematics", "Create quadratic equations questions")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Title != "Quadratic Equations - Practice Set" {
		t.Errorf("expected Title 'Quadratic Equations - Practice Set', got %q", resp.Title)
	}
	if resp.Instructions != "Solve all questions and show your working." {
		t.Errorf("expected Instructions 'Solve all questions and show your working.', got %q", resp.Instructions)
	}
	if len(resp.Questions) != 1 {
		t.Fatalf("expected 1 question, got %d", len(resp.Questions))
	}
	if resp.Questions[0].Question != "Solve x² - 5x + 6 = 0" {
		t.Errorf("expected question text 'Solve x² - 5x + 6 = 0', got %q", resp.Questions[0].Question)
	}
	if resp.TotalMarks != 2 {
		t.Errorf("expected total marks 2, got %d", resp.TotalMarks)
	}
}

func TestAIServiceImpl_GenerateAssignment_MarkdownFences(t *testing.T) {
	mockClient := &mockLLMClient{
		generateFn: func(ctx context.Context, systemPrompt, userPrompt string) (string, error) {
			return "```json\n" + `{
				"title": "Fractions Quiz",
				"instructions": "Answer all questions.",
				"questions": [
					{"question": "What is 1/2 + 1/4?", "type": "short_answer", "marks": 5},
					{"question": "What is 3/4 - 1/4?", "type": "short_answer", "marks": 5}
				],
				"total_marks": 10
			}` + "\n```", nil
		},
	}
	service := NewAIService(mockClient)

	resp, err := service.GenerateAssignment(context.Background(), "Grade 5", "Math", "Fractions questions")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Title != "Fractions Quiz" {
		t.Errorf("expected Title 'Fractions Quiz', got %q", resp.Title)
	}
	if len(resp.Questions) != 2 {
		t.Fatalf("expected 2 questions, got %d", len(resp.Questions))
	}
	if resp.TotalMarks != 10 {
		t.Errorf("expected TotalMarks 10, got %d", resp.TotalMarks)
	}
}

func TestAIServiceImpl_GenerateAssignment_LLMError(t *testing.T) {
	mockClient := &mockLLMClient{
		generateFn: func(ctx context.Context, systemPrompt, userPrompt string) (string, error) {
			return "", errors.New("rate limit exceeded")
		},
	}
	service := NewAIService(mockClient)

	_, err := service.GenerateAssignment(context.Background(), "Grade 10", "Math", "Quadratic equations")
	if err == nil {
		t.Fatal("expected error on LLM failure, got nil")
	}
	if !strings.Contains(err.Error(), "ai generation failed") {
		t.Errorf("expected error message to contain 'ai generation failed', got %q", err.Error())
	}
}

func TestAIServiceImpl_GenerateAssignment_InvalidJSON(t *testing.T) {
	mockClient := &mockLLMClient{
		generateFn: func(ctx context.Context, systemPrompt, userPrompt string) (string, error) {
			return "Here is your assignment: not valid json", nil
		},
	}
	service := NewAIService(mockClient)

	_, err := service.GenerateAssignment(context.Background(), "Grade 10", "Math", "Quadratic equations")
	if err == nil {
		t.Fatal("expected error on invalid JSON, got nil")
	}
	if !strings.Contains(err.Error(), "failed to parse AI response JSON") {
		t.Errorf("expected error message to contain 'failed to parse AI response JSON', got %q", err.Error())
	}
}

func TestAIServiceImpl_GenerateAssignment_Validation(t *testing.T) {
	mockClient := &mockLLMClient{}
	service := NewAIService(mockClient)

	tests := []struct {
		name      string
		className string
		subject   string
		prompt    string
	}{
		{"empty class", "", "Mathematics", "10 questions"},
		{"empty subject", "Grade 10", "", "10 questions"},
		{"empty prompt", "Grade 10", "Mathematics", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := service.GenerateAssignment(context.Background(), tt.className, tt.subject, tt.prompt)
			if err == nil {
				t.Error("expected error for invalid input, got nil")
			}
		})
	}
}

func TestAIServiceImpl_NilLLMClient(t *testing.T) {
	service := NewAIService(nil)
	_, err := service.GenerateAssignment(context.Background(), "Grade 10", "Math", "Prompt")
	if err == nil {
		t.Error("expected error when LLM client is nil")
	}
}

func TestPrompts(t *testing.T) {
	sys := SystemPrompt()
	if !strings.Contains(sys, "JSON") || !strings.Contains(sys, "questions") {
		t.Errorf("expected system prompt to contain JSON schema instructions, got %s", sys)
	}

	usr := UserPrompt("Grade 10-A", "Physics", "Thermodynamics laws")
	if !strings.Contains(usr, "Grade 10-A") || !strings.Contains(usr, "Physics") || !strings.Contains(usr, "Thermodynamics laws") {
		t.Errorf("expected user prompt to contain class, subject, and topic, got %s", usr)
	}
}
