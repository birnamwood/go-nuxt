package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// User comment
type User struct {
	Mail     string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// Signup の説明
func Signup(c echo.Context) error {
	user := new(User)
	//user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	//model.CreateUser(user)

	return c.String(http.StatusOK, user.Name)
	//return c.JSON(http.StatusOK, user)

}
