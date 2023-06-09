package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/amirfakhrullah/go-chat/pkg/cli"
	"github.com/amirfakhrullah/go-chat/pkg/helpers"
	_ "github.com/joho/godotenv/autoload"
	openai "github.com/sashabaranov/go-openai"
)

var apiKeyName = "OPEN_AI_API_KEY"

func main() {
	var apiKey string
	apiKey = os.Getenv(apiKeyName)

	if len(apiKey) == 0 {
		var err error
		apiKey, err = cli.GetApiKey(apiKeyName)
		helpers.HandlePanic(err)
	}

	c := openai.NewClient(apiKey)
	ctx := context.Background()

	var input string
	var inputErr error
	input, inputErr = cli.GetQuestion(true)
	helpers.HandlePanic(inputErr)

	for input != ":q!" {
		req := openai.ChatCompletionRequest{
			Model:     openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: input,
				},
			},
			Stream: true,
		}

		stream, err := c.CreateChatCompletionStream(ctx, req)
		helpers.HandlePanic(err)

		defer stream.Close()

		fmt.Print("Response: ")
		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				fmt.Print("\n")
				break
			}
			helpers.HandlePanic(err)

			fmt.Printf(response.Choices[0].Delta.Content)
		}

		input, inputErr = cli.GetQuestion(false)
		helpers.HandlePanic(inputErr)
	}
}
