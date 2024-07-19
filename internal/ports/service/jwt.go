package service

import (
	"api/autentiacion/internal/domain"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	GetSecret() string
	CreateToken(user *domain.UserLogin) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	ParseToken(tokenString string) (*domain.UserLogin, error)
}
