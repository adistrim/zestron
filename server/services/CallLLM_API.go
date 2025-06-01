package services

import (
	"context"
	"os"
	"zestron-server/models"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func CallLLM_API(prompt string, history []models.Message) (string, error) {
	systemPrompt := `You are a helpful assistant. Respond to the user's query directly and completely using the fewest possible words.
Your response must contain only the information necessary to fully address the user's request. Do not add any extra details,
explanations, conversational elements, or creative content. Be accurate, concise, and professional while prioritizing extreme brevity.`

	messages := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(systemPrompt),
	}
	for _, msg := range history {
		switch msg.Role {
		case "user":
			messages = append(messages, openai.UserMessage(msg.Content))
		case "assistant":
			messages = append(messages, openai.AssistantMessage(msg.Content))
		}
	}
	messages = append(messages, openai.UserMessage(prompt))

	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("KEY")),
		option.WithBaseURL("https://api.deepseek.com"),
	)

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: messages,
		Model:    "deepseek-chat",
	})
	if err != nil {
		return "", err
	}
	return chatCompletion.Choices[0].Message.Content, nil
}
