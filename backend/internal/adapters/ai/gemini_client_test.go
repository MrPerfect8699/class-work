package ai

import (
	"context"
	"testing"
)

func TestNewGeminiClient_EmptyAPIKey(t *testing.T) {
	_, err := NewGeminiClient(context.Background(), "", "gemini-2.5-flash")
	if err == nil {
		t.Error("expected error when api key is empty, got nil")
	}
}

func TestNewGeminiClient_DefaultModel(t *testing.T) {
	client, err := NewGeminiClient(context.Background(), "test-key", "")
	if err != nil {
		t.Fatalf("unexpected error creating client: %v", err)
	}
	if client.model != "gemini-3.6-flash" {
		t.Errorf("expected default model 'gemini-3.6-flash', got %q", client.model)
	}
}

func TestGeminiClient_NilClientGenerate(t *testing.T) {
	client := &GeminiClient{client: nil}
	_, err := client.Generate(context.Background(), "system", "user")
	if err == nil {
		t.Error("expected error when genai.Client is nil, got nil")
	}
}
