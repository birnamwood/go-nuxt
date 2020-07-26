package api

import (
	"net/http"

	"../../model"
	"github.com/labstack/echo"
)

// User comment

// Signup の説明
func Signup(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	user.Create()

	return c.JSON(http.StatusOK, user)

}
