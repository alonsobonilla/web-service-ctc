package handlers

import (
	authports "ctcwebservice/core/auth/ports"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService authports.AuthService
}

func NewAuthHandler(authService authports.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	err := h.authService.Login(email, password)
	if err != nil {
		c.JSON(400, gin.H{})
	}

	c.JSON(200, gin.H{"message": "Hello wordl!"})
}
