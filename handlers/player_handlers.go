package handlers

import (
	playerdomain "ctcwebservice/core/player/domain"
	playerports "ctcwebservice/core/player/ports"
	"net/http"

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

	playerId, err := h.playerService.Register(&player)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"player_id": playerId}, "message": "Registrado correctamente"})
}
