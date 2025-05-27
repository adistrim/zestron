package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"zestron-server/models"
	"zestron-server/services"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func WebSocketHandler(chatManager *services.ChatManager, allowedOrigins []string) gin.HandlerFunc {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			if origin == "" {
				return false
			}

			for _, allowed := range allowedOrigins {
				if origin == strings.TrimSpace(allowed) {
					return true
				}
			}

			return false
		},
	}

	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Printf("Error upgrading to WebSocket: %v", err)
			return
		}
		defer conn.Close()

		for {
			messageType, msg, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("Error reading message: %v", err)
				}
				break
			}

			var request models.WebSocketRequest
			if err := json.Unmarshal(msg, &request); err != nil {
				log.Printf("Error parsing message: %v", err)
				sendErrorResponse(conn, "", "Invalid request format", messageType)
				continue
			}

			var chatSession *models.ChatSession
			var exists bool

			if request.ChatID != "" {
				chatSession, exists = chatManager.GetSession(request.ChatID)
				if !exists {
					chatSession = chatManager.CreateSession()
				}
			} else {
				chatSession = chatManager.CreateSession()
			}

			userMessage := models.Message{
				Role:    "user",
				Content: request.Prompt,
			}

			chatManager.AddMessage(chatSession.ID, userMessage)

			response, err := services.CallLLM_API(request.Prompt, chatSession.Messages[:len(chatSession.Messages)-1])
			if err != nil {
				log.Printf("Error calling LLM API: %v", err)
				sendErrorResponse(conn, chatSession.ID, "Failed to generate response", messageType)
				continue
			}

			assistantMessage := models.Message{
				Role:    "assistant",
				Content: response,
			}

			chatManager.AddMessage(chatSession.ID, assistantMessage)

			wsResponse := models.WebSocketResponse{
				ChatID:   chatSession.ID,
				Message:  assistantMessage,
				Complete: true,
			}

			responseJSON, err := json.Marshal(wsResponse)
			if err != nil {
				log.Printf("Error marshaling response: %v", err)
				continue
			}

			if err := conn.WriteMessage(messageType, responseJSON); err != nil {
				log.Printf("Error writing message: %v", err)
				break
			}
		}
	}
}

func sendErrorResponse(conn *websocket.Conn, chatID, errorMsg string, messageType int) {
	wsResponse := models.WebSocketResponse{
		ChatID:   chatID,
		Error:    errorMsg,
		Complete: true,
	}

	responseJSON, err := json.Marshal(wsResponse)
	if err != nil {
		log.Printf("Error marshaling error response: %v", err)
		return
	}

	if err := conn.WriteMessage(messageType, responseJSON); err != nil {
		log.Printf("Error sending error response: %v", err)
	}
}
