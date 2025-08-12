package handlers

import (
	"chatroom/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MessageHandler struct {
	DB *gorm.DB
}

type CreateMessageRequest struct {
	Content string `json:"content" binding:"required"`
}

func (h *MessageHandler) CreateMessage(c *gin.Context) {
	var req CreateMessageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	message := models.Message{
		Content:  req.Content,
		Username: c.GetString("username"),
		UserID:   c.GetUint("userID"),
	}

	if err := h.DB.Create(&message).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create message",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Message created",
	})
}

func (h *MessageHandler) GetMessages(c *gin.Context) {
	var messages []models.Message

	if err := h.DB.Find(&messages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get messages",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messages": messages,
	})
}
