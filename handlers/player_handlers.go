package handlers

import (
	playerdomain "ctcwebservice/core/player/domain"
	playerports "ctcwebservice/core/player/ports"

	"github.com/gin-gonic/gin"
)

type PlayerHandler struct {
	playerService playerports.PlayerService
}

func NewPlayerHandler(playerService playerports.PlayerService) *PlayerHandler {
	return &PlayerHandler{
		playerService: playerService,
	}
}
func (h *PlayerHandler) Register(c *gin.Context) {
	var player playerdomain.Player
	c.BindJSON(&player)

	err := h.playerService.Register(&player)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	if err == nil {
		c.JSON(200, gin.H{"message": "Hello wordl!"})
	}
}
