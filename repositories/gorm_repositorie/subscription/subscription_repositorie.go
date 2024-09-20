package subscriptionrepositorie

import (
	subscriptiondomain "ctcwebservice/core/subscription/domain"
	subscriptionports "ctcwebservice/core/subscription/ports"
	"ctcwebservice/repositories"
	"errors"
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

func (r *SubscriptionRepository) Register(s *subscriptiondomain.Subscription) (*subscriptiondomain.Subscription, error) {
	result := r.db.Create(s)

	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return nil, repositories.ErrDuplicatedKey
	}
	return s, nil
}

func (r *SubscriptionRepository) ChangeState(s *subscriptiondomain.Subscription) (*subscriptiondomain.Subscription, error) {
	return nil, repositories.ErrRecordNotFound
}

func (r *SubscriptionRepository) UpdateLastSearch(t time.Time) (*subscriptiondomain.Subscription, error) {
	return nil, repositories.ErrRecordNotFound
}
