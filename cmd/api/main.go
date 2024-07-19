package main

import (
	"api/autentiacion/cmd/api/handlers/user"
	"api/autentiacion/cmd/api/middleware"
	"api/autentiacion/internal/repositories/mongo"
	userRpo "api/autentiacion/internal/repositories/mongo/user"
	"api/autentiacion/internal/services/auth"
	userSrv "api/autentiacion/internal/services/user"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// TODO: implementar el servidor de la API
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cliente, err := mongo.ConnectClient(os.Getenv("MONGO_URL"))

	if err != nil {
		log.Fatal(err)
	}

	// var userRepository user.UserRepository =

	userRepository := userRpo.UserRepository{
		Client:     cliente,
		Collection: cliente.Database(os.Getenv("MONGO_DATABASE")).Collection("users"),
	}

	userService := userSrv.Service{
		UserRepository: userRepository,
	}

	userHandler := user.Handler{
		UserService: userService,
		JwtService:  *auth.NewJwtSecret(),
	}

	ginEngine := gin.Default()
	ginEngine.POST("/api/users/register", userHandler.Register)
	// ginEngine.Use(middleware.Passport(userHandler))
	ginEngine.GET("/", middleware.Passport(userHandler), user.GetRoot)
	ginEngine.POST("/api/users/create", middleware.Passport(userHandler), userHandler.Create)
	ginEngine.POST("/api/users/login", userHandler.Login)

	log.Fatalln(ginEngine.Run(":8001"))

}
