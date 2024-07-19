package auth

import "os"

type AuthService struct {
	Secret string
}

func NewJwtSecret() *AuthService {
	return &AuthService{
		Secret: os.Getenv("JWT_SECRET"),
	}
}
