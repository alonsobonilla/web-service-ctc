package sportdomain

type Sport struct {
	SportID     uint8   `gorm:"primaryKey" json:"sport_id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Image       *string `json:"image"`
}
