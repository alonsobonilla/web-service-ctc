package playerports

import (
	playerdomain "ctcwebservice/core/player/domain"

	"github.com/gin-gonic/gin"
)

type PlayerRepository interface {
	Create(*playerdomain.Player) (uint, error)
}

type PlayerService interface {
	Register(*playerdomain.Player) (uint, error)
}

type PlayerHandler interface {
	Register(*gin.Context)
}
