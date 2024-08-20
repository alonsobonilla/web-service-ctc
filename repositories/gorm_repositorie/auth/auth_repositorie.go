package authgorm

import (
	playerdomain "ctcwebservice/core/player/domain"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) Login(email string) (playerdomain.Player, error) {
	return playerdomain.Player{}, nil
}
