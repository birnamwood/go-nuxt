package api

import (
	"net/http"
	"strconv"

	"../../database"
	"../../model"
	"github.com/labstack/echo"
)

// ShowUser の説明
func ShowUser(c echo.Context) error {
	var id int
	id, _ = strconv.Atoi(c.Param("user_id"))

	user := new(model.User)
	database.Init(user)
	user.FindByID(id)

	return c.JSON(http.StatusOK, user)
}
