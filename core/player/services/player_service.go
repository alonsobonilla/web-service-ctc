package playerservice

import (
	playerdomain "ctcwebservice/core/player/domain"
	playerports "ctcwebservice/core/player/ports"
)

type PlayerService struct {
	playerRepository playerports.PlayerRepository
}

func NewPlayerService(playerRepository playerports.PlayerRepository) *PlayerService {
	return &PlayerService{
		playerRepository: playerRepository,
	}
}

func (s *PlayerService) Register(player *playerdomain.Player) error {
	var err error

	err = player.HashPassword()
	if err != nil {
		return err
	}
	err = s.playerRepository.Create(player)
	return err
}
