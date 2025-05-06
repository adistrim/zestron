package models

type GenerateRequest struct {
	Prompt string `json:"prompt" binding:"required"`
}

type GenerateResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}
