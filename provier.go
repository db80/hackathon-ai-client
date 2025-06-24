package main

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/bedrock"
	"github.com/tmc/langchaingo/llms/googleai"
	"log"
	"os"
)

type GetProvider interface {
	get() llms.Model
}
type BedRockProvider struct{}

func (b *BedRockProvider) get() llms.Model {
	llm, err := bedrock.New(
		bedrock.WithModel("anthropic.claude-3-7-sonnet-20250219-v1:0"),
	)
	if err != nil {
		log.Fatalf("failed to create Bedrock LLM: %v", err)
	}
	return llm
}

type GeminiProvider struct {
	APIKey string
}

func (g *GeminiProvider) get() llms.Model {
	llm, err := googleai.New(
		context.Background(),
		googleai.WithAPIKey(g.APIKey),
		googleai.WithDefaultModel("gemini-2.5-flash"),
	)
	if err != nil {
		log.Fatalf("failed to create Gemini LLM: %v", err)
	}
	return llm
}

func NewProviderFromEnv() (GetProvider, error) {
	provider := os.Getenv("LLM_PROVIDER")
	switch provider {
	case "bedrock":
		return &BedRockProvider{}, nil
	case "gemini":
		apiKey := os.Getenv("GOOGLE_API_KEY")
		if apiKey == "" {
			return nil, fmt.Errorf("GOOGLE_API_KEY environment variable not set")
		}
		return &GeminiProvider{APIKey: apiKey}, nil
	default:
		return nil, fmt.Errorf("unknown LLM_PROVIDER: %s", provider)
	}
}
