package role

import (
	"api/autentiacion/internal/ports/service"
)

type Handler struct {
	RoleService service.RoleService
}
