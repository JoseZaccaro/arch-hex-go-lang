package user

import (
	"api/autentiacion/internal/ports/service"
	"api/autentiacion/internal/services/auth"
)

type Handler struct {
	UserService service.UserService
	JwtService  auth.AuthService
}
