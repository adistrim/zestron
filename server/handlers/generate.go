package handlers

import (
	"zestron-server/models"
	"zestron-server/services"

	"github.com/gin-gonic/gin"
)

func GenerateHandler(c *gin.Context) {
	var req models.GenerateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	response, err := services.CallLLM_API(req.Prompt)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to call LLM API"})
		return
	}

	c.JSON(200, gin.H{"response": response})
}
