package playerdomain

import (
	sportdomain "ctcwebservice/core/sport/domain"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// type Player struct {
// 	PlayerID     uint
// 	UserName     string    `json:"user_name"`
// 	LastName     string    `json:"last_name"`
// 	FirstName    string    `json:"first_name"`
// 	Password     string    `json:"password"`
// 	PhoneNumber  string    `json:"phone_number"`
// 	Email        string    `json:"email"`
// 	Gender       string    `json:"gender"`
// 	BirthDate    time.Time `json:"birth_date"`
// 	PlayerImage  string    `json:"player_image"`
// 	CreatedAt    time.Time
// 	UpdatedAt    time.Time
// 	PlayerSports []sportdomain.Sport
// }

type Player struct {
	PlayerID     uint   `gorm:"primaryKey"`
	FirstName    string `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName     string `gorm:"type:varchar(100);not null" json:"last_name"`
	Password     string `gorm:"type:varchar(255);not null" json:"password"`
	PhoneNumber  string `gorm:"type:char(9);unique;not null" json:"phone_number"`
	Email        string `gorm:"type:varchar(255);unique;not null" json:"email"`
	Gender       string `gorm:"type:char(1);not null;" json:"gender"`
	BirthDate    string `gorm:"type:date" json:"birth_date"`
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
