package ollama

import (
	"aiocvgen/internal/config"
	"aiocvgen/internal/utils"
	"context"
	"strings"

	"github.com/ollama/ollama/api"
)

// Raw generate function that handles everything
func generate(systemPrompt string, prompt string, conf config.Conf) string {
	// Client setup
	var resp = ""
	Stage = StageLoading

	// Generation!
	client.Generate(context.Background(), &api.GenerateRequest{
		Model:  conf.OllamaModel,
		Prompt: prompt,
		System: systemPrompt,
		Stream: utils.Bool(conf.OllamaStream),
		Think:  &api.ThinkValue{Value: conf.OllamaThink},
		Options: map[string]any{
			"temperature": conf.OllamaTemperature,
		},
	}, func(gr api.GenerateResponse) error {
		if Stage == StageIdle {
			return nil
		}
		lthk := len(gr.Thinking)
		lrsp := len(gr.Response)
		if lthk > 0 {
			Stage = StageThinking
		} else if lrsp > 0 {
			Stage = StageWriting
			resp += gr.Response
		}
		return nil
	})

	if Stage == StageIdle {
		return ""
	}
	resp = strings.TrimSpace(resp)
	Stage = StageIdle
	return resp
}

// Uses default (empty) system prompt for general use
func Generate(prompt string) string {
	c := config.Get()
	c.OllamaThink = false
	x := generate("", prompt, c)
	return x
}

// Uses non-default system prompt that helps in generating HTML CVs
func GenerateResume(prompt string, conf config.Conf) string {
	return generate(
		`You are an expert specializing in generating resumes. Match the language of the provided instructions. Your goal is to produce a well-formatted, professional resume in a single HTML file. You will get information you need to strictly follow:
		- Candidate info - fully-detailed person description, like names, contact, life, experience and everything else
		- Job offer - a copy-pasted job announcement content listed on employer's website
		- HTML layout/look options - how exactly HTML-based resume should be stylized
		- Notes - other general but still important information
		Sections are separated by markdown-formatted titles starting with #`,
		prompt, conf,
	)
}
