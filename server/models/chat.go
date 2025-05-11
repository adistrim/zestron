package models

import (
	"time"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatSession struct {
	ID       string    `json:"id"`
	Messages []Message `json:"messages"`
	Created  time.Time `json:"created"`
}

type WebSocketRequest struct {
	ChatID string `json:"chatId,omitempty"`
	Prompt string `json:"prompt"`
}

type WebSocketResponse struct {
	ChatID   string  `json:"chatId"`
	Message  Message `json:"message"`
	Error    string  `json:"error,omitempty"`
	Complete bool    `json:"complete"`
}
