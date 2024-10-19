package feature

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type llmFeature struct {
	model *genai.GenerativeModel
}

func NewLLMFeature(apiKey string) *llmFeature {
	ctx := context.Background()

	llmClient, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
		return nil
	}

	modelName := "gemini-1.5-flash"
	model := llmClient.GenerativeModel(modelName)

	log.Printf("[LLM] %s has been started...", modelName)

	return &llmFeature{model: model}
}

func (f llmFeature) GenerateHelpText(word string) string {
	prompt := fmt.Sprintf("I do not know spanish. Translate the \"%s\". List some exmaple sentences or phrases with this word. Generate compact answer in 2 very small paragraphs. Make examples in spanish as bullet (-) points", word)

	resp, err := f.model.GenerateContent(context.Background(), genai.Text(prompt))
	if err != nil {
		log.Print(err)
		return ""
	}

	result := fmt.Sprintf("%s", getStringResponse(resp))

	return result
}

func getStringResponse(resp *genai.GenerateContentResponse) string {
	var result string

	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				result += fmt.Sprintln(part)
			}
		}
	}

	return result
}
