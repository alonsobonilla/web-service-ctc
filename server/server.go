package server

import (
	authports "ctcwebservice/core/auth/ports"
	playerports "ctcwebservice/core/player/ports"

	"github.com/gin-gonic/gin"
)

type Server struct {
	authHandler   authports.AuthHandler
	playerHandler playerports.PlayerHandler
}

func NewServer(authHandler authports.AuthHandler, playerHandler playerports.PlayerHandler) *Server {
	return &Server{
		authHandler:   authHandler,
		playerHandler: playerHandler,
	}
}

func (s *Server) Initialize() {

	r := gin.Default()

	group := r.Group("/api/v1")

	group.GET("/helloworld", s.authHandler.Login)
	group.POST("/player", s.playerHandler.Register)

	r.Run("localhost:8080")
}
