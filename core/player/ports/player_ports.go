package playerports

import (
	playerdomain "ctcwebservice/core/player/domain"

	"github.com/gin-gonic/gin"
)

type PlayerRepository interface {
	Create(player *playerdomain.Player) (uint, error)
}

type PlayerService interface {
	Register(player *playerdomain.Player) (uint, error)
}

type PlayerHandler interface {
	Register(c *gin.Context)
}
