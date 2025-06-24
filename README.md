# Hackathon AI Client

A Go-based AI client that supports multiple LLM providers including AWS Bedrock and Google Gemini. This application demonstrates text generation and image analysis capabilities using the LangChain Go library.

## Features

- Support for multiple LLM providers (AWS Bedrock, Google Gemini)
- Text generation capabilities
- Image analysis with tag generation
- Provider selection via environment variables

## Prerequisites

- Go 1.24.2 or higher
- AWS credentials configured (for Bedrock provider)
- Google API credentials (for Gemini provider)

## Installation

```bash
git clone <repository-url>
cd hackathon-ai-client
go mod download
```

## Environment Variables

The following environment variables are required to run the application:

### Required

- `LLM_PROVIDER` - Specifies which LLM provider to use. Valid values:
  - `bedrock` - Uses AWS Bedrock with Claude 3.7 Sonnet model
  - `gemini` - Uses Google Gemini 2.5 Flash model

### Provider-Specific

#### For Bedrock Provider
- AWS credentials must be configured via:
  - `AWS_ACCESS_KEY_ID` - Your AWS access key
  - `AWS_SECRET_ACCESS_KEY` - Your AWS secret key
  - `AWS_REGION` - AWS region (`eu-west-2`)

#### For Gemini Provider
- `GOOGLE_API_KEY` - Your Google AI API key (required when `LLM_PROVIDER=gemini`)

## Usage

### Running with Bedrock

```bash
export LLM_PROVIDER=bedrock
export AWS_REGION=us-east-1
# Ensure AWS credentials are configured
go run .
```

### Running with Gemini

```bash
export LLM_PROVIDER=gemini
export GOOGLE_API_KEY=your-google-api-key-here
go run .
```

## Example Output

The application will:
1. Generate a joke about Brazilian football players
2. Analyze the `hox_ny.png` image and return JSON-formatted tags
3. Display token usage information

## Project Structure

- `main.go` - Main application entry point
- `provier.go` - LLM provider implementations
- `hox_ny.png` - Sample image for analysis
- `go.mod` - Go module dependencies

## Dependencies

This project uses:
- [LangChain Go](https://github.com/tmc/langchaingo) - LLM integration library
- AWS SDK for Go v2 - For Bedrock integration
- Google Generative AI Go - For Gemini integration

## Error Handling

The application will exit with an error if:
- `LLM_PROVIDER` is not set or contains an invalid value
- `GOOGLE_API_KEY` is not set when using the Gemini provider
- AWS credentials are not properly configured when using the Bedrock provider
- The image file `hox_ny.png` is not foundq