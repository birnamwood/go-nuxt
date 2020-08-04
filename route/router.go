package route

import (
	"github.com/birnamwood/go-nuxt/api"
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

	v1 := e.Group("/api/v1")
	{
		v1.GET("/user/:user_id", api.ShowUser)
	}
	v2 := e.Group("/api/v2")
	{
		v2.Use(middleware.JWTWithConfig(handler.Config)) // /api 下はJWTの認証が必要
	}

	e.Logger.Fatal(e.Start(":" + c.GetString("server.port")))
	return nil
}
