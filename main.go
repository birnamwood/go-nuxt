package main

import (
	"./handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routing
	e.GET("/:clinicname/login", handler.LoginPage())
	e.GET("/:clinicname/main", handler.MainPage())

	// Setup
	e.Start(":8080")

}
