package route

import (
	"../api/v1"
	"../config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//Init router
func Init() *echo.Echo {
	c := config.GetConfig()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routing
	v1 := e.Group("/api/v1")
	{
		v1.POST("/signup", api.Signup)
		v1.GET("/user/:user_id", api.ShowUser)
	}

	e.Logger.Fatal(e.Start(":" + c.GetString("server.port")))
	return nil
}
