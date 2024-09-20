package main

import (
	authservice "ctcwebservice/core/auth/service"
	playerservice "ctcwebservice/core/player/services"
	subscriptionservice "ctcwebservice/core/subscription/services"
	"ctcwebservice/handlers"
	gormrepositorie "ctcwebservice/repositories/gorm_repositorie"
	authgorm "ctcwebservice/repositories/gorm_repositorie/auth"
	playergorm "ctcwebservice/repositories/gorm_repositorie/player"
	subscriptionrepositorie "ctcwebservice/repositories/gorm_repositorie/subscription"
	"ctcwebservice/server"
	"ctcwebservice/utils"

	"github.com/gin-gonic/gin/binding"
)

func main() {
	binding.Validator = new(utils.CustomValidator)

	playerRepository := playergorm.NewPlayerRepository(gormrepositorie.Db)
	playerService := playerservice.NewPlayerService(playerRepository)
	playerHandler := handlers.NewPlayerHandler(playerService)

	authRepository := authgorm.NewAuthRepository(gormrepositorie.Db)
	authService := authservice.NewAuthService(authRepository)
	authHandlers := handlers.NewAuthHandler(authService)

	subscriptionRepository := subscriptionrepositorie.NewSubscriptionRepository(gormrepositorie.Db)
	subscriptionService := subscriptionservice.NewSubscriptionService(subscriptionRepository)
	subscriptionHandlers := handlers.NewSubscriptionHandler(subscriptionService)
	server := server.NewServer(authHandlers, playerHandler, subscriptionHandlers)

	server.Initialize()
}
