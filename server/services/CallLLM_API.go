package services

import (
	"os"
	"zestron-server/models"

	"github.com/go-resty/resty/v2"
)

func CallLLM_API(prompt string) (string, error) {
	client := resty.New()

	var result models.GenerateResponse

	systemPrompt := `You are a helpful assistant. Respond to the user's query directly and completely using the fewest possible words.
	Your response must contain only the information necessary to fully address the user's request. Do not add any extra details,
	explanations, conversational elements, or creative content. Be accurate, concise, and professional while prioritizing extreme brevity.`

	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+os.Getenv("KEY")).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]any{
			"model": "deepseek-chat",
			"messages": []any{
				map[string]string{
					"role":    "system",
					"content": systemPrompt,
				},
				map[string]string{
					"role":    "user",
					"content": prompt,
				},
			},
		}).
		SetResult(&result).
		Post("https://api.deepseek.com/chat/completions")

	if err != nil || resp.IsError() {
		return "", err
	}

	if len(result.Choices) > 0 {
		return result.Choices[0].Message.Content, nil
	}
	return "", nil
}
