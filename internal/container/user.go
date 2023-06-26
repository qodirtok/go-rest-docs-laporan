package container

import (
	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/user/handlers"
	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/user/repositories"
	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/user/services"
	"gorm.io/gorm"
)

func BuildContainerUser(db *gorm.DB) *handlers.UserHandler {
	userRepository := repositories.NewUserRepositories(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	return userHandler
}
