package subscriptionports

import (
	subscriptiondomain "ctcwebservice/core/subscription/domain"
	"time"

	"github.com/gin-gonic/gin"
)

type SubscriptionRepository interface {
	Register(*subscriptiondomain.Subscription) error
	ChangeState(*subscriptiondomain.Subscription) error
	UpdateLastSearch(time.Time) error
}

type SubscriptionService interface {
	Register(subscriptiondomain.Subscription) error
	ChangeState(subscriptiondomain.Subscription) error
	UpdateLastSearch(time.Time) error
}

type SubscriptionHandler interface {
	Register(*gin.Context)
	ChangeState(*gin.Context)
	UpdateLastSearch(*gin.Context)
}
