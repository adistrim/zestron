package services

import (
	"sync"
	"time"
	"zestron-server/models"

	"github.com/google/uuid"
)

type ChatManager struct {
	sessions map[string]*models.ChatSession
	mutex    sync.RWMutex
}

func NewChatManager() *ChatManager {
	return &ChatManager{
		sessions: make(map[string]*models.ChatSession),
	}
}

func (cm *ChatManager) CreateSession() *models.ChatSession {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	id := uuid.New().String()
	session := &models.ChatSession{
		ID:       id,
		Messages: []models.Message{},
		Created:  time.Now(),
	}

	cm.sessions[id] = session
	return session
}

func (cm *ChatManager) GetSession(id string) (*models.ChatSession, bool) {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	session, exists := cm.sessions[id]
	return session, exists
}

func (cm *ChatManager) AddMessage(chatID string, message models.Message) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	if session, exists := cm.sessions[chatID]; exists {
		session.Messages = append(session.Messages, message)
	}
}
