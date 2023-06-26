package container

import (
	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/login/handlers"
	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/login/repositories"
	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/login/services"
	"gorm.io/gorm"
)

func BuildLoginContainer(db *gorm.DB) *handlers.LoginHandler {
	LoginRepository := repositories.NewLoginRepositories(db)
	LoginService := services.NewLoginService(LoginRepository)
	LoginHandler := handlers.NewLoginHandler(LoginService)
	return LoginHandler
}
