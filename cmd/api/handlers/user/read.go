package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetRoot(c *gin.Context) {
	users, _ := h.UserService.ReadAll(c)

	for _, user := range users {
		fmt.Println(user.RoleID)
		if roleID, ok := user.RoleID.(string); ok {
			user.RoleID = roleID
		}
	}

	c.JSON(200, gin.H{"message": "Hello World", "users": users})
}
