package role

import "github.com/gin-gonic/gin"

func (h Handler) ReadAll(c *gin.Context) {
	roles, _ := h.RoleService.ReadAll(c)
	c.JSON(200, gin.H{"message": "All roles", "Roles": roles})
}
