package handlers

import (
	subscriptiondomain "ctcwebservice/core/subscription/domain"
	subscriptionports "ctcwebservice/core/subscription/ports"
	handlers "ctcwebservice/handlers/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SubscriptionHandler struct {
	service subscriptionports.SubscriptionService
}

func NewSubscriptionHandler(service subscriptionports.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{
		service: service,
	}
}

func (h *SubscriptionHandler) Register(c *gin.Context) {
	var err error
	var s *subscriptiondomain.Subscription
	var subscription subscriptiondomain.Subscription

	if err = c.ShouldBindJSON(&subscription); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": nil, "error": err.Error()})
		return
	}

	if s, err = h.service.Register(&subscription); err != nil {
		c.JSON(handlers.Mapped[err], gin.H{"data": nil, "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": s, "mensaje": "Subscription created"})
}

func (h *SubscriptionHandler) ChangeState(c *gin.Context) {

}

func (h *SubscriptionHandler) UpdateLastSearch(c *gin.Context) {

}
