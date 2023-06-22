package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/user/models"
	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/user/services"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omiempty"`
}

type UserHandler struct {
	userService services.ServicesUser
}

func NewUserHandler(userService services.ServicesUser) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) PostUserHandler(c *gin.Context) {
	var input models.UserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorMessages := make([]string, len(validationErrors))
			for i, e := range validationErrors {
				errorMessages[i] = fmt.Sprintf("error on field %s, condition: %s", e.Field(), e.ActualTag())
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errorMessages,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	user, err := h.userService.Create(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": map[string]interface{}{"message": "User successfully saved", "data": convertRespondUser(user)},
	})
}

func (h *UserHandler) GetUserHandler(c *gin.Context) {
	user, err := h.userService.FindAll()

	if err != nil {
		c.JSON(http.StatusOK, Response{Success: true, Error: "failed To get Data"})
		return
	}

	var usersResponse []models.UserResponse

	for _, u := range user {
		userResponse := convertRespondUser(u)
		usersResponse = append(usersResponse, userResponse)
	}

	c.JSON(http.StatusOK, Response{Success: true, Data: usersResponse})
}

func convertRespondUser(u models.User) models.UserResponse {
	return models.UserResponse{
		ID:        int(u.ID),
		Name:      u.Name,
		Email:     u.Email,
		Username:  u.Username,
		CreatedAt: u.CreatedAt.String(),
		UpdatedAt: u.UpdatedAt.String(),
	}
}
