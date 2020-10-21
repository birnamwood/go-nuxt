package api_v1

import (
	"net/http"
	"strconv"

	"github.com/birnamwood/go-nuxt/model"
	"github.com/birnamwood/go-nuxt/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// GetCurrentUser の説明
func GetCurrentUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	mail := claims["email"].(string)
	var currentUser = repository.FindUserByEmail(mail)
	return c.JSON(http.StatusOK, currentUser)
}

// ShowUser の説明
func ShowUser(c echo.Context) error {
	var id int
	id, _ = strconv.Atoi(c.Param("user_id"))
	var user model.User
	user = repository.FindUserByID(id)

	return c.JSON(http.StatusOK, user)
}

func ShowUsers(c echo.Context) error {
	var users []model.User
	users = repository.FindAllUsers()

	return c.JSON(http.StatusOK, users)
}
