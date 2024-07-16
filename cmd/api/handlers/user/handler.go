package user

import "api/autentiacion/internal/ports/service"

type Handler struct {
	UserService service.UserService
}
