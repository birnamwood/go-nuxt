package handler

import (
	"github.com/labstack/echo"
	"github.com/birnamwood/go-nuxt/model"
)

// Signup の説明
func Signup(c echo.Context) error {
	user:= new(model.User)
	if err := c:Bind(u); err != nil {
		return err
	}

}

// Signin の説明
func Signin() echo.HandlerFunc {

}
