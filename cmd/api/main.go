package main

import (
	"api/autentiacion/internal/repositories/mongo"
	repoRoleMongo "api/autentiacion/internal/repositories/mongo/role"
	"api/autentiacion/internal/repositories/mysql"
	repoUserMySql "api/autentiacion/internal/repositories/mysql/user"
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

	// Conectar a MongoDB
	cliente, err := mongo.ConnectClient(os.Getenv("MONGO_URL"))
	if err != nil {
		log.Fatal(err)
	}

	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")

	// Conectar a MySql
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")" + "/autenticacion" + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := mysql.Connect(dsn)
	if err != nil {
		log.Fatal(err)
	}

	// var userRepository user.UserRepository =
	roleRepository := repoRoleMongo.RoleRepository{
		Client:     cliente,
		Collection: cliente.Database(os.Getenv("MONGO_DATABASE")).Collection("roles"),
	}
	// userRepository := repoUserMongo.UserRepository{
	// 	Client:     cliente,
	// 	Collection: cliente.Database(os.Getenv("MONGO_DATABASE")).Collection("users"),
	// }
	userRepositorySql := repoUserMySql.UserRepository{
		DB: db,
	}

	userService := userSrv.Service{
		UserRepository: userRepositorySql,
		RoleRepository: roleRepository,
	}

	userHandler := user.Handler{
		UserService: userService,
		JwtService:  *auth.NewJwtSecret(),
	}

	ginEngine := gin.Default()
	ginEngine.POST("/api/users/register", userHandler.Register)
	// ginEngine.Use(middleware.Passport(userHandler))
	ginEngine.GET("/", userHandler.GetRoot)
	ginEngine.POST("/api/users/create", userHandler.Create)
	ginEngine.POST("/api/users/login", userHandler.Login)
	ginEngine.PUT("/api/users/:id", userHandler.UpdateUser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001" // Valor predeterminado si la variable de entorno no está configurada
	}

	log.Fatalln(ginEngine.Run(":" + port))

}
