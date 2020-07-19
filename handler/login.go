package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// LoginPage の説明
func LoginPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		clinicName := c.Param("clinicname")
		return c.String(http.StatusOK, clinicName+"ログインページ")
	}
}
