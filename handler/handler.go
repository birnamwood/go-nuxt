package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// MainPage の説明
func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		clinicName := c.Param("clinicname")
		jsonMap := map[string]string{
			"クリニック名": clinicName,
			"患者名":    "テスト患者",
		}

		return c.JSON(http.StatusOK, jsonMap)
	}
}
