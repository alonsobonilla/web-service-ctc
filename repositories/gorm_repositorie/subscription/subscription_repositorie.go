package subscriptionrepositorie

import (
	subscriptiondomain "ctcwebservice/core/subscription/domain"
	subscriptionports "ctcwebservice/core/subscription/ports"
	"ctcwebservice/repositories"
	"time"

	"gorm.io/gorm"
)

type SubscriptionRepository struct {
	db *gorm.DB
}

var _ subscriptionports.SubscriptionRepository = (*SubscriptionRepository)(nil)

func NewSubscriptionRepository(db *gorm.DB) *SubscriptionRepository {
	return &SubscriptionRepository{
		db: db,
	}
}

func (r *SubscriptionRepository) Register(s *subscriptiondomain.Subscription) error {
	return repositories.ErrRecordNotFound
}

func (r *SubscriptionRepository) ChangeState(s *subscriptiondomain.Subscription) error {
	return repositories.ErrRecordNotFound
}

func (r *SubscriptionRepository) UpdateLastSearch(t time.Time) error {
	return repositories.ErrRecordNotFound
}
