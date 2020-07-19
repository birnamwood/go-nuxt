package main

import (
	"./handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func newRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routing
	e.POST("/signup", handler.Signup())
	e.POST("/login", handler.Login())

	return e
}
