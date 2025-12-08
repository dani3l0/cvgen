package ollama

import (
	"aiocvgen/internal/config"
	"aiocvgen/internal/utils"
	"context"
	"strings"

	"github.com/ollama/ollama/api"
)

// Raw generate function that handles everything
func generate(systemPrompt string, prompt string, think bool) string {
	// Client setup
	var resp = ""
	Stage = Stages.Loading

	// Generation!
	client.Generate(context.Background(), &api.GenerateRequest{
		Model:  config.Get().OllamaModel,
		Prompt: prompt,
		System: systemPrompt,
		Stream: utils.Bool(config.Get().OllamaStream),
		Think:  &api.ThinkValue{Value: !think},
		Options: map[string]any{
			"temperature": config.Get().OllamaTemperature,
		},
	}, func(gr api.GenerateResponse) error {
		lthk := len(gr.Thinking)
		lrsp := len(gr.Response)
		if lthk > 0 {
			Stage = Stages.Thinking
		} else if lrsp > 0 {
			Stage = Stages.Writing
			resp += gr.Response
		}
		return nil
	})

	resp = strings.TrimSpace(resp)
	Stage = Stages.Done
	return resp
}

// Uses default (empty) system prompt for general use
func Generate(prompt string) string {
	x := generate("", prompt, false)
	Stage = Stages.Idle
	return x
}

// Uses non-default system prompt that helps in generating HTML CVs
func GenerateResume(prompt string) string {
	return generate(
		"You are an expert resume writer specializing in generating resumes in HTML and CSS. Your goal is to produce a well-formatted, professional resume in a single HTML file. Maintain a professional tone and match the language of the provided instructions.",
		prompt,
		config.Get().OllamaThink,
	)
}
