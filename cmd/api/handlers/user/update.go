package user

import (
	"api/autentiacion/internal/domain"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h Handler) UpdateUser(c *gin.Context) {

	var userParams domain.UserCreateParams

	if err := c.BindJSON(&userParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(userParams)
	user, err := h.UserService.Update(c, c.Param("id"), userParams)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Update", "user": user})
}
