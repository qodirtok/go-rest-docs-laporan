package helper

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	login "github.com/qodirtok/go-rest-docs-laporan/internal/pkg/login/models"
	"gorm.io/gorm"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

var Database *gorm.DB

func GenerateJWT(login login.Login) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       login.ID,
		"username": login.Username,
		"iat":      time.Now().Unix(),
		"eat":      time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})

	return token.SignedString(privateKey)
}

func CurrentUser(c *gin.Context) (login.Login, error) {

	err := ValidateJWT(c)
	if err != nil {
		return login.Login{}, err
	}

	token, _ := getToken(c)
	claims, _ := token.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	var userDb login.Login
	err = Database.Where("id =?", userId).Find(&userDb).Error
	if err != nil {
		return login.Login{}, err
	}

	return userDb, nil
}

// func findByUsername

func getToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

func getTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}

func ValidateJWT(context *gin.Context) error {
	token, err := getToken(context)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}
