package main

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
	"github.com/xh3b4sd/tracer"
)

func main() {
	//
}

func call(inp string) (string, error) {
	cli := openai.NewClient("TODO")
	ctx := context.Background()

	req := openai.CompletionRequest{
		Model:     openai.GPT3Ada,
		MaxTokens: 1024,
		Prompt:    inp,
	}

	resp, err := cli.CreateCompletion(ctx, req)
	if err != nil {
		return "", tracer.Mask(err)
	}

	return resp.Choices[0].Text, nil
}
