package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/user/models"
	"github.com/qodirtok/go-rest-docs-laporan/internal/pkg/user/services"
	"gorm.io/gorm"
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

	// userAuth, err := helper.CurrentUser(c)

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err,
	// 	})
	// 	return
	// }

	// if input.Username != userAuth.Username{
	// 	input.Username = userAuth.Username
	// }

	user, err := h.userService.Create(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	// c.JSON(http.StatusCreated, gin.H{
	// 	"result": map[string]interface{}{"message": "User successfully saved", "data": convertRespondUser(user)},
	// })
	c.JSON(http.StatusCreated, Response{Success: true, Data: convertRespondUser(user)})
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

func (h *UserHandler) GetUserByIdHandler(c *gin.Context) {

	idUserStr := c.Param("id")

	log.Println("id = ", idUserStr)
	idUser, err := strconv.Atoi(idUserStr)

	log.Println("id = ", idUser)
	if err != nil {
		c.JSON(http.StatusOK, Response{Success: false, Error: "ID User is Invalid"})
		return
	}

	user, err := h.userService.FindOne(idUser)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response := Response{
				Success: false,
				Error:   "User not found",
			}
			c.JSON(http.StatusNotFound, response)
			return
		}
		c.JSON(http.StatusOK, Response{Success: true, Error: "failed To get Data"})
		return
	}

	c.JSON(http.StatusOK, Response{Success: true, Data: convertRespondUser(user)})
}

func (h *UserHandler) UserDeleteHandler(c *gin.Context) {
	idUserStr := c.Param("id")

	log.Println("id = ", idUserStr)
	idUser, err := strconv.Atoi(idUserStr)

	log.Println("id = ", idUser)
	if err != nil {
		c.JSON(http.StatusOK, Response{Success: false, Error: "ID User is Invalid"})
		return
	}

	user, err := h.userService.Delete(idUser)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response := Response{
				Success: false,
				Error:   "User not found",
			}
			c.JSON(http.StatusNotFound, response)
			return
		}
		c.JSON(http.StatusOK, Response{Success: true, Error: "failed To get Data"})
		return
	}

	c.JSON(http.StatusOK, Response{Success: true, Data: convertRespondUser(user)})
}

func (h *UserHandler) UserUpdateHandler(c *gin.Context) {
	idUserStr := c.Param("id")

	log.Println("id = ", idUserStr)
	idUser, _ := strconv.Atoi(idUserStr)

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

	user, err := h.userService.UpdateUser(idUser, input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"result": map[string]interface{}{"message": "User successfully saved", "data": convertRespondUser(user)},
	// })
	c.JSON(http.StatusOK, Response{Success: true, Data: convertRespondUser(user)})
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
