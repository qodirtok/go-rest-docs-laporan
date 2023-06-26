package services

import (
	"html"
	"strings"

	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/user/models"
	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/user/repositories"
)

type ServicesUser interface {
	Create(user models.UserInput) (models.User, error)
	FindAll() ([]models.User, error)
	FindOne(ID int) (models.User, error)
	Delete(ID int) (models.User, error)
	UpdateUser(ID int, user models.UserInput) (models.User, error)
}

type userService struct {
	repository repositories.UserRepositories
}

func NewUserService(repository repositories.UserRepositories) ServicesUser {
	return &userService{repository}
}

func (s *userService) Create(userInput models.UserInput) (models.User, error) {

	// passwordHash, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)

	// if err != nil {
	// 	return models.User{}, err
	// }

	user := models.User{
		Name:     userInput.Name,
		Username: html.EscapeString(strings.TrimSpace(userInput.Username)),
		// Password: string(passwordHash),
		Email: userInput.Email,
	}

	newUser, err := s.repository.Create(user)
	return newUser, err
}

func (s *userService) FindAll() ([]models.User, error) {
	user, err := s.repository.FindAll()
	return user, err
}

func (s *userService) FindOne(ID int) (models.User, error) {
	user, err := s.repository.FindOne(ID)
	return user, err
}

func (s *userService) Delete(ID int) (models.User, error) {
	user, err := s.repository.Delete(ID)
	return user, err
}

func (s *userService) UpdateUser(ID int, input models.UserInput) (models.User, error) {
	user, err := s.repository.FindOne(ID)

	user.Name = input.Name
	user.Email = input.Email
	user.Username = html.EscapeString(strings.TrimSpace(input.Username))

	newUser, err := s.repository.UpdateUser(user)

	return newUser, err
}
