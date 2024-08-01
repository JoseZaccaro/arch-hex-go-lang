package user

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetRoot(c *gin.Context) {
	users, _ := h.UserService.ReadAll(c)
	for _, user := range users {
		log.Println(user.ID)
	}
	c.JSON(200, gin.H{"message": "Hello World", "users": users})
}
