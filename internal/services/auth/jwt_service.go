package auth

import (
	"api/autentiacion/internal/domain"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (j AuthService) GetSecret() string {
	return j.Secret
}

func (j AuthService) CreateToken(user *domain.UserLogin) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      user.ID,
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", fmt.Errorf("error signing token: %v", err)
	}

	return tokenString, nil
}

func (j AuthService) ValidateToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (j AuthService) ParseToken(token *jwt.Token) (*domain.UserLogin, error) {

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	user := domain.UserLogin{
		ID:       claims["sub"].(string),
		Username: claims["username"].(string),
		Email:    claims["email"].(string),
	}

	return &user, nil
}
