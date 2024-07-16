package user

import (
	"api/autentiacion/internal/domain"

	"github.com/gin-gonic/gin"
)

func (h Handler) Create(c *gin.Context) {
	//? => Traducir el request
	var userParams domain.UserCreateParams
	if err := c.BindJSON(&userParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//? => Hacer validaciones
	// ! Validar password
	// ! Validar email
	// ! Validar username

	//? => Utilizar el service
	user, err := h.UserService.Create(c, userParams)

	if err != nil {
		c.JSON(500, gin.H{"error": "Error creating user"})
	}

	//? => Traducir el response
	c.JSON(200, gin.H{"message": "Created", "userID": user.ID})
}
