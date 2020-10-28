package user_api

import (
	"net/http"
	"strconv"

	"github.com/birnamwood/go-nuxt/pkg/logger"
	"github.com/birnamwood/go-nuxt/pkg/model"
	"github.com/birnamwood/go-nuxt/pkg/repository/userRepository"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func Hello(c echo.Context) error {
	logger.NewLogger().Info("API呼び出し")
	return c.String(http.StatusOK, "Hello, World!")
}

// GetCurrentUser の説明
func GetCurrentUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	mail := claims["email"].(string)
	var currentUser = userRepository.FindUserByEmail(mail)
	return c.JSON(http.StatusOK, currentUser)
}

// ShowUser の説明
func ShowUser(c echo.Context) error {
	var id int
	id, _ = strconv.Atoi(c.Param("user_id"))
	var user model.User
	user = userRepository.FindUserByID(id)

	return c.JSON(http.StatusOK, user)
}

func ShowUsers(c echo.Context) error {
	var users []model.User
	users = userRepository.FindAllUsers()

	return c.JSON(http.StatusOK, users)
}
