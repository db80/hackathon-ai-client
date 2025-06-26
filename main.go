package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tmc/langchaingo/llms"
)

func main() {
	provider, err := NewLLMProvider()
	if err != nil {
		log.Fatalf("failed to create provider: %v", err)
	}
	fmt.Printf("using provider: %T\n", provider)

	llm := provider.get()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Choose an action:")
	fmt.Println("1. Talk to LLM")
	fmt.Println("2. Extract tags from image")
	fmt.Print("Enter your choice (1 or 2): ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		talkToLLM(llm)
	case "2":
		extractImageTags(llm)
	default:
		fmt.Println("Invalid choice")
	}
}

func talkToLLM(llm llms.Model) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your prompt: ")
	prompt, _ := reader.ReadString('\n')
	prompt = strings.TrimSpace(prompt)

	response, err := llms.GenerateFromSinglePrompt(context.Background(), llm, prompt)
	if err != nil {
		log.Fatalf("failed to generate response: %v", err)
	}

	fmt.Println("\nResponse:")
	fmt.Println(response)
}

var imageFileName = "hox_ny.png"

func extractImageTags(llm llms.Model) {
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
					llms.TextPart("Analyze the provided image and generate a list of relevant tags. The response must be a JSON object with a single key named 'tags', which contains an array of strings representing the tags."),
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

	fmt.Println("\nImage Tags:")
	fmt.Println(choices[0].Content)
}