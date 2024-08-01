package main

import (
	RoleHandler "api/autentiacion/cmd/api/handlers/role"
	UserHandler "api/autentiacion/cmd/api/handlers/user"
	"api/autentiacion/cmd/api/middleware"
	mongoPkg "api/autentiacion/internal/repositories/mongo"
	repoRoleMongo "api/autentiacion/internal/repositories/mongo/role"
	mysqlPkg "api/autentiacion/internal/repositories/mysql"
	repoUserMySql "api/autentiacion/internal/repositories/mysql/user"
	authPkg "api/autentiacion/internal/services/auth"
	roleSrv "api/autentiacion/internal/services/role"
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
	mongoUrl := os.Getenv("MONGO_URL")
	// Conectar a MongoDB
	cliente, err := mongoPkg.ConnectClient(mongoUrl)
	if err != nil {
		log.Fatal(err)
	}

	mysql_host := os.Getenv("MYSQL_HOST")
	mysql_port := os.Getenv("MYSQL_PORT")
	mysql_user := os.Getenv("MYSQL_USER")
	mysql_password := os.Getenv("MYSQL_PASSWORD")

	// Conectar a MySql

	dsn := mysql_user + ":" + mysql_password + "@tcp(" + mysql_host + ":" + mysql_port + ")" + "/autenticacion" + "?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:admin@tcp(localhost:3306)/autenticacion?charset=utf8mb4&parseTime=True&loc=Local"

	db, errSql := mysqlPkg.Connect(dsn)
	if errSql != nil {
		log.Fatal(errSql)
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

	userHandler := UserHandler.Handler{
		UserService: userService,
		JwtService:  *authPkg.NewJwtSecret(),
	}

	roleService := roleSrv.Service{
		RoleRepository: roleRepository,
	}
	roleHandler := RoleHandler.Handler{
		RoleService: roleService,
	}

	ginEngine := gin.Default()
	ginEngine.POST("/api/users/register", userHandler.Register)
	ginEngine.POST("/api/users/login", userHandler.Login)
	// ginEngine.Use(middleware.Passport(userHandler))
	ginEngine.GET("/", middleware.Passport(userHandler), userHandler.GetRoot)
	ginEngine.POST("/api/users/create", userHandler.Create)
	ginEngine.PUT("/api/users/:id", userHandler.UpdateUser)
	ginEngine.POST("/api/role", roleHandler.Create)
	ginEngine.GET("/api/role", roleHandler.ReadAll)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001" // Valor predeterminado si la variable de entorno no está configurada
	}

	log.Fatalln(ginEngine.Run(":" + port))

}
