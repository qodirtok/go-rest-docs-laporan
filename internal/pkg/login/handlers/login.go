package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/login/helper"
	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/login/models"
	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/login/services"
	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omiempty"`
}

// type LoginHandler interface {
// }

type LoginHandler struct {
	service services.LoginService
}

func NewLoginHandler(service services.LoginService) *LoginHandler {
	return &LoginHandler{service}
}

func (h *LoginHandler) PostLoginCreate(c *gin.Context) {
	var input models.AuthentificationInput

	if err := c.ShouldBindJSON(&input); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorMessages := make([]string, len(validationErrors))
			for i, e := range validationErrors {
				errorMessages[i] = fmt.Sprintf("error on field %s, condition %s ", e.Field(), e.ActualTag())
			}
			c.JSON(http.StatusBadRequest, gin.H{"Success": true, "Error": errorMessages})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"Success": true, "Error": "Internal Server Error"})
		return
	}

	register, err := h.service.Save(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Success": true, "Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": true, "Data": register})
}

func (h *LoginHandler) Login(c *gin.Context) {
	var input models.AuthentificationInput
	fmt.Sprint(input)

	if err := c.ShouldBindJSON(&input); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorMessages := make([]string, len(validationErrors))
			for i, e := range validationErrors {
				errorMessages[i] = fmt.Sprintf("error on field %s, condition %s ", e.Field(), e.ActualTag())
			}
			c.JSON(http.StatusBadRequest, gin.H{"Success": true, "Error": errorMessages})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"Success": true, "Error": "Internal Server Error"})
		return
	}

	// fmt.Sprint(input)
	login, err := h.service.FindByUsername(input.Username)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "Error": err.Error()})
		return
	}

	err = ValidatePassword(login, input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Success": false, "Error": err.Error()})
		return
	}

	jwt, err := helper.GenerateJWT(login)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": true, "token": jwt})
}

func ValidatePassword(p models.Login, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(password))
}

func response(l models.Login) *models.AuthResponse {
	return &models.AuthResponse{
		Username:  l.Username,
		CreatedAt: l.CreatedAt.String(),
		UpdatedAt: l.UpdatedAt.String(),
	}
}
