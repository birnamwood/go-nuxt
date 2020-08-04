package api

import (
	"net/http"
	"strconv"

	"github.com/birnamwood/go-nuxt/model"
	"github.com/birnamwood/go-nuxt/repository"
	"github.com/labstack/echo"
)

// ShowUser の説明
func ShowUser(c echo.Context) error {
	var id int
	id, _ = strconv.Atoi(c.Param("user_id"))

	var user model.User
	user = repository.FindUserByID(id)

	return c.JSON(http.StatusOK, user)
}
