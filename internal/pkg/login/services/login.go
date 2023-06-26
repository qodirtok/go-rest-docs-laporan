package services

import (
	"html"
	"strings"

	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/login/models"
	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/login/repositories"
	"golang.org/x/crypto/bcrypt"
)

type LoginService interface {
	Save(loginInput models.AuthentificationInput) (models.Login, error)
	FindByUsername(username string) (models.Login, error)
}

type service struct {
	repository repositories.LoginRepositories
}

func NewLoginService(repository repositories.LoginRepositories) LoginService {
	return &service{repository}
}

func (s *service) Save(input models.AuthentificationInput) (models.Login, error) {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		return models.Login{}, err
	}

	auth := models.Login{
		Username: html.EscapeString(strings.TrimSpace(input.Username)),
		Password: string(passwordHash),
	}

	// fmt.Println(auth)

	newLogin, err := s.repository.Save(auth)

	return newLogin, err
}

func (s *service) FindByUsername(username string) (models.Login, error) {
	var login models.Login

	login, err := s.repository.FindByUsername(username)
	return login, err
}

// func (s *service) ValidatePassword(password string) (models.Login, error) {
// 	// var login models.Login
// 	return bcrypt.CompareHashAndPassword([]byte(s.repository.), []byte(password))
// }
