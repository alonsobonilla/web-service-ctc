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

func (r *PlayerRepository) Create(player *playerdomain.Player) error {

	var result *gorm.DB
	result = r.db.Create(player)

	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		result = r.db.First(player, "email = ?", player.Email)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return repositories.ErrDuplicatedKeyPhoneNumber
		}

		return repositories.ErrDuplicatedKeyEmail
	}
	return nil
}
