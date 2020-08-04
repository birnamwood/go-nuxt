package route

import (
	"github.com/birnamwood/go-nuxt/api/v1"
	"github.com/birnamwood/go-nuxt/config"
	"github.com/birnamwood/go-nuxt/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//Init router
func Init() *echo.Echo {
	c := config.GetConfig()
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routing
	e.POST("/signup", handler.Signup)
	e.POST("/login", handler.Login)
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.POST("", handler.Restricted())

	v1 := e.Group("/api/v1")
	{
		v1.GET("/user/:user_id", api.ShowUser)
	}
	v2 := e.Group("/api/v2")
	{
		v2.Use(middleware.JWT([]byte("secret")))
	}

	e.Logger.Fatal(e.Start(":" + c.GetString("server.port")))
	return nil
}
