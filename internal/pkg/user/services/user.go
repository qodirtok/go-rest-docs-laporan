package services

import (
	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/user/models"
	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/user/repositories"
	"golang.org/x/crypto/bcrypt"
)

type ServicesUser interface {
	Create(user models.UserInput) (models.User, error)
	FindAll() ([]models.User, error)
}

type userService struct {
	repository repositories.UserRepositories
}

func NewUserService(repository repositories.UserRepositories) ServicesUser {
	return &userService{repository}
}

func (s *userService) Create(userInput models.UserInput) (models.User, error) {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)

	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Name:     userInput.Name,
		Username: userInput.Username,
		Password: string(passwordHash),
		Email:    userInput.Email,
	}

	newUser, err := s.repository.Create(user)
	return newUser, err
}

func (s *userService) FindAll() ([]models.User, error) {
	user, err := s.repository.FindAll()
	return user, err
}
