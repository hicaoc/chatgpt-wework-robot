package main

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

var openaiconfig openai.ClientConfig
var openaiclient *openai.Client

func chat(msglist []openai.ChatCompletionMessage) (string, error) {

	resp, err := openaiclient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: msglist,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	//fmt.Println(resp.Choices[0].Message.Content)
	return resp.Choices[0].Message.Content, nil
}
