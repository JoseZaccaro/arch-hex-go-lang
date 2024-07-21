package user

import (
	"api/autentiacion/internal/domain"
	"log"

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
	id, err := h.UserService.Create(c, userParams)

	if err != nil {
		c.JSON(500, gin.H{"error": "Error creating user"})
		return
	}

	//? => Traducir el response
	c.JSON(200, gin.H{"message": "Created", "userID": id})
}

func (h Handler) Register(c *gin.Context) {
	var userParams domain.UserCreateParams

	if err := c.BindJSON(&userParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserService.Register(c, userParams)

	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(200, gin.H{"message": "Register", "user": user})
}

func (h Handler) Login(c *gin.Context) {
	var userParams domain.UserCreateParams
	if err := c.BindJSON(&userParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserService.Login(c, userParams.Email, userParams.PasswordHash)

	if err != nil {
		log.Println(err)
		if err.Error() == "username or password incorrect" {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"error": "Error login user"})
		return
	}

	token, err := h.JwtService.CreateToken(user)

	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "Error creating token"})
		return
	}

	// c.SetCookie("token", token, 3600, "/", "localhost", false, true)

	c.JSON(200, gin.H{"message": "Login", "token": token})

}
