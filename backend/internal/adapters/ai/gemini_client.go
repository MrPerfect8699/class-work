package ai

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/genai"
)

// GeminiClient implements ports.LLMClient using the Google Gen AI Go SDK.
type GeminiClient struct {
	client *genai.Client
	model  string
}

// NewGeminiClient creates a new GeminiClient instance.
func NewGeminiClient(ctx context.Context, apiKey, model string) (*GeminiClient, error) {
	if apiKey == "" {
		return nil, errors.New("gemini api key is required")
	}

	if model == "" {
		model = "gemini-3.6-flash"
	}

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create gemini client: %w", err)
	}

	return &GeminiClient{
		client: client,
		model:  model,
	}, nil
}

// Generate sends prompts to the Gemini model and returns the string response.
func (g *GeminiClient) Generate(ctx context.Context, systemPrompt, userPrompt string) (string, error) {
	if g.client == nil {
		return "", errors.New("gemini client is not initialized")
	}

	config := &genai.GenerateContentConfig{
		SystemInstruction: &genai.Content{
			Parts: []*genai.Part{
				{Text: systemPrompt},
			},
		},
		ResponseMIMEType: "application/json",
	}

	resp, err := g.client.Models.GenerateContent(ctx, g.model, genai.Text(userPrompt), config)
	if err != nil {
		return "", fmt.Errorf("gemini generate content failed: %w", err)
	}

	text := resp.Text()
	if text == "" && len(resp.Candidates) > 0 && resp.Candidates[0].Content != nil {
		for _, part := range resp.Candidates[0].Content.Parts {
			if part.Text != "" {
				text += part.Text
			}
		}
	}

	if text == "" {
		return "", errors.New("empty response received from gemini")
	}

	return text, nil
}
