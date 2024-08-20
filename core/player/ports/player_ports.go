package playerports

import (
	playerdomain "ctcwebservice/core/player/domain"

	"github.com/gin-gonic/gin"
)

type PlayerRepository interface {
	Create(player *playerdomain.Player) error
}

type PlayerService interface {
	Register(player *playerdomain.Player) error
}

type PlayerHandler interface {
	Register(c *gin.Context)
}
