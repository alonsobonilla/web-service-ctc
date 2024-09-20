package subscriptionservice

import (
	subscriptiondomain "ctcwebservice/core/subscription/domain"
	subscriptionports "ctcwebservice/core/subscription/ports"
	"time"
)

type SubscriptionService struct {
	repository subscriptionports.SubscriptionRepository
}

func NewSubscriptionService(r subscriptionports.SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{
		repository: r,
	}
}

func (ser *SubscriptionService) Register(s *subscriptiondomain.Subscription) (*subscriptiondomain.Subscription, error) {
	return ser.repository.Register(s)
}

func (ser *SubscriptionService) ChangeState(s *subscriptiondomain.Subscription) (*subscriptiondomain.Subscription, error) {
	return ser.repository.ChangeState(s)

}

func (ser *SubscriptionService) UpdateLastSearch(t time.Time) (*subscriptiondomain.Subscription, error) {
	return ser.repository.UpdateLastSearch(t)
}
