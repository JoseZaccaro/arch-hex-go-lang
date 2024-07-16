package main

import (
	"api/autentiacion/cmd/api/handlers/user"
	userSrv "api/autentiacion/internal/services/user"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// TODO: implementar el servidor de la API
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	userService := userSrv.NewService()

	userHandler := user.Handler{
		UserService: userService,
	}

	ginEngine := gin.Default()
	ginEngine.GET("/", user.GetRoot)
	ginEngine.POST("/api/users/create", userHandler.Create)

	log.Fatalln(ginEngine.Run(":8001"))
}
