package cmd

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/qodirtok/go-rest-docs-laporan/internal/config"
	"github.com/qodirtok/go-rest-docs-laporan/internal/container"
	"github.com/qodirtok/go-rest-docs-laporan/internal/middleware"
)

func Getenv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error Loading .env file")
	}
}

func Cmd() {

	Getenv()
	db := config.ConnectionDB()

	// Migration(db)
	Migration(db)

	user := container.BuildContainerUser(db)
	login := container.BuildLoginContainer(db)

	rest := gin.Default()
	rest.Use(middleware.LoggerMiddleware)

	v1 := rest.Group("/v1/auth")
	v1.POST("/register", login.PostLoginCreate)
	v1.POST("/login", login.Login)
	auth := rest.Group("/v1/api")
	auth.Use(middleware.JWTAuthMiddleware())
	auth.POST("/user", user.PostUserHandler)
	auth.GET("/user", user.GetUserHandler)
	auth.GET("/user/:id", user.GetUserByIdHandler)
	auth.POST("/user/:id", user.UserUpdateHandler)
	auth.DELETE("/user/:id", user.UserDeleteHandler)

	err := rest.Run(":8000")

	if err != nil {
		log.Fatal("Failed to start server : ", err)
	}
}
