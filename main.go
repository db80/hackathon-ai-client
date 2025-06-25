package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tmc/langchaingo/llms"
)

func main() {
	provider, err := NewLLMProvider()
	if err != nil {
		log.Fatalf("failed to create provider: %v", err)
	}
	fmt.Printf("using provider: %T\n", provider)

	llm := provider.get()
	prompt := "Tell me a joke about brazilian football players."

	response, err := llms.GenerateFromSinglePrompt(context.Background(), llm, prompt)
	if err != nil {
		fmt.Errorf("failed to generate response: %w", err)
	}

	fmt.Println(response)
	fmt.Println(images(llm))
}

var imageFileName = "hox_ny.png"

func images(llm llms.Model) string {
	imgData, err := os.ReadFile(imageFileName)
	if err != nil {
		log.Fatalf("failed to read image file: %v", err)
	}

	resp, err := llm.GenerateContent(
		context.Background(),
		[]llms.MessageContent{
			{
				Role: llms.ChatMessageTypeHuman,
				Parts: []llms.ContentPart{
					llms.BinaryPart("image/png", imgData),
					llms.TextPart("Give me a list ot tags for this image. - response should be in JSON format with a single key 'tags' and an array of strings as value."),
				},
			},
		},
		llms.WithMaxTokens(3000),
		llms.WithTemperature(0.1),
		llms.WithTopP(1.0),
		llms.WithTopK(100),
		llms.WithJSONMode(),
	)
	if err != nil {
		log.Fatal(err)
	}
	choices := resp.Choices
	if len(choices) < 1 {
		log.Fatal("empty response from model")
	}

	return choices[0].Content
}
