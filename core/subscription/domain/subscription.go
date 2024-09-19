package subscriptiondomain

import "time"

type Subscription struct {
	PlayerID   uint      `gorm:"primaryKey" json:"player_id" binding:"required,number"`
	SportID    uint      `gorm:"primaryKey" json:"sport_id" binding:"required,number"`
	Active     bool      `gorm:"type:boolean;not null;default:false" json:"active"`
	LastSearch time.Time `gorm:"type:timestamp"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
