package playergorm

import (
	playerdomain "ctcwebservice/core/player/domain"
	"ctcwebservice/repositories"
	"errors"

	"gorm.io/gorm"
)

type PlayerRepository struct {
	db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) *PlayerRepository {
	return &PlayerRepository{
		db: db,
	}
}

func (r *PlayerRepository) GetPlayerByEmail(email string) (playerdomain.Player, error) {
	return playerdomain.Player{}, nil
}

func (r *PlayerRepository) Create(player *playerdomain.Player) (uint, error) {

	var result *gorm.DB
	result = r.db.Create(player)

	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		result = r.db.First(player, "email = ? and phone_number = ? ", player.Email, player.PhoneNumber)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			result = r.db.First(player, "email = ?", player.Email)
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return 0, repositories.ErrDuplicatedKeyPhoneNumber
			}
			return 0, repositories.ErrDuplicatedKeyEmail
		}

		return 0, repositories.ErrDuplicatedKeyEmailPhoneNumber
	}
	return player.PlayerID, nil
}
