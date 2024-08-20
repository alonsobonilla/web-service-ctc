package authservice

import (
	authports "ctcwebservice/core/auth/ports"
)

type AuthService struct {
	authRepository authports.AuthRepository
}

func NewAuthService(authRepository authports.AuthRepository) *AuthService {
	return &AuthService{
		authRepository: authRepository,
	}
}

func (service *AuthService) Login(email string, password string) error {
	player, err := service.authRepository.Login(email)

	if err != nil {
		return err
	}

	if player.Password == password {
		return nil
	}
	return nil
}
