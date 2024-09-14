package playerdomain

import (
	sportdomain "ctcwebservice/core/sport/domain"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Player struct {
	PlayerID     uint   `gorm:"primaryKey"`
	FirstName    string `gorm:"type:varchar(100);not null" json:"first_name" binding:"required"`
	LastName     string `gorm:"type:varchar(100);not null" json:"last_name" binding:"required"`
	Password     string `gorm:"type:varchar(255);not null" json:"password" binding:"required"`
	PhoneNumber  string `gorm:"type:char(9);unique;not null" json:"phone_number" binding:"required,len=9"`
	Email        string `gorm:"type:varchar(255);unique;not null" json:"email" binding:"required,email"`
	Gender       string `gorm:"type:char(1);not null;" json:"gender" binding:"required"`
	BirthDate    string `gorm:"type:date" json:"birth_date" binding:"required"`
	PlayerImage  string `json:"player_image"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	PlayerSports []sportdomain.Sport `gorm:"many2many:player_sports;foreignKey:PlayerID;joinForeignKey:PlayerID;References:SportID;joinReferences:SportID"`
}

func (p *Player) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	p.Password = string(bytes)
	return nil
}

func (p *Player) ComparePasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(password))
	return err == nil
}
