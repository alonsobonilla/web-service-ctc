package authports

import (
	playerdomain "ctcwebservice/core/player/domain"

	"github.com/gin-gonic/gin"
)

type AuthRepository interface {
	Login(email string) (playerdomain.Player, error)
}

type AuthHandler interface {
	Login(c *gin.Context)
}

type AuthService interface {
	Login(email string, password string) error
}
