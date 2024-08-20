package playersportdomain

import "time"

type PlayerSport struct {
	PlayerID  uint `gorm:"primaryKey;autoIncrement:false"`
	SportID   uint `gorm:"primaryKey;autoIncrement:false"`
	State     bool
	CreatedAt time.Time
}
