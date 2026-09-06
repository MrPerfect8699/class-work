package ports

import "context"

// LLMClient defines the interface for interacting with an LLM provider.
type LLMClient interface {
	Generate(ctx context.Context, systemPrompt, userPrompt string) (string, error)
}
