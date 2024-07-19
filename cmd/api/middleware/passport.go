package middleware

import (
	"api/autentiacion/cmd/api/handlers/user"
	"strings"

	"github.com/gin-gonic/gin"
)

func Passport(h user.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		// log.Println("Authorization: ", authorization)
		if authorization == "" || !strings.HasPrefix(authorization, "Bearer ") {
			UnauthorizeAuth(c)
			return
		}
		token := strings.Split(authorization, " ")[1]
		// log.Println("Token: ", token)
		jwt, err := h.JwtService.ValidateToken(token)
		// log.Println("JWT: ", jwt)
		if err != nil {
			UnauthorizeAuth(c)
			return
		}

		user, errParsing := h.JwtService.ParseToken(jwt)
		// log.Println("User: ", user)
		if errParsing != nil {
			UnauthorizeAuth(c)
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

func UnauthorizeAuth(c *gin.Context) {
	c.JSON(401, gin.H{"message": "Unauthorized"})
	c.Abort()
}
