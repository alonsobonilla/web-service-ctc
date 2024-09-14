package main

import (
	authservice "ctcwebservice/core/auth/service"
	playerservice "ctcwebservice/core/player/services"
	"ctcwebservice/handlers"
	gormrepositorie "ctcwebservice/repositories/gorm_repositorie"
	authgorm "ctcwebservice/repositories/gorm_repositorie/auth"
	playergorm "ctcwebservice/repositories/gorm_repositorie/player"
	"ctcwebservice/server"
	"ctcwebservice/utils"

	"github.com/gin-gonic/gin/binding"
)

func main() {
	binding.Validator = new(utils.Validator)

	playerRepository := playergorm.NewPlayerRepository(gormrepositorie.Db)
	playerService := playerservice.NewPlayerService(playerRepository)
	playerHandler := handlers.NewPlayerHandler(playerService)

	authRepository := authgorm.NewAuthRepository(gormrepositorie.Db)
	authService := authservice.NewAuthService(authRepository)
	authHandlers := handlers.NewAuthHandler(authService)

	server := server.NewServer(authHandlers, playerHandler)

	server.Initialize()
}
