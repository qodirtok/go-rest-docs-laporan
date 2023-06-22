package cmd

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/qodirtok/go-rest-docs-laporan/internal/config"
	"github.com/qodirtok/go-rest-docs-laporan/internal/middleware"
	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/user/container"
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

	migration(db)

	user := container.BuildContainerUser(db)

	rest := gin.Default()
	rest.Use(middleware.LoggerMiddleware)

	v1 := rest.Group("/v1")
	v1.POST("/user", user.PostUserHandler)
	v1.GET("/user", user.GetUserHandler)

	err := rest.Run(":8000")

	if err != nil {
		log.Fatal("Failed to start server : ", err)
	}
}
