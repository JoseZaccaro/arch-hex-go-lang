package role

import (
	"api/autentiacion/internal/domain"

	"github.com/gin-gonic/gin"
)

func (h Handler) Create(c *gin.Context) {
	//? => Traducir el request
	var roleParams domain.CreateRole
	if err := c.BindJSON(&roleParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//? => Utilizar el service
	id, err := h.RoleService.Create(c, roleParams)

	if err != nil {
		c.JSON(500, gin.H{"error": "Error creating role"})
		return
	}

	//? => Traducir el response
	c.JSON(200, gin.H{"message": "Created", "RoleID": id})
}
