package subscriptiondomain

import "time"

type Subscription struct {
	PlayerID   uint      `gorm:"primaryKey" json:"player_id"`
	SportID    uint      `gorm:"primaryKey" json:"sport_id"`
	Active     bool      `gorm:"type:boolean;not null" json:"active"`
	LastSearch time.Time `gorm:"type:timestamp;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
