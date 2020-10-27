package route

import (
	"github.com/birnamwood/go-nuxt/config"
	"github.com/birnamwood/go-nuxt/pkg/api/v1/clinic_api"
	"github.com/birnamwood/go-nuxt/pkg/api/v1/user_api"
	"github.com/birnamwood/go-nuxt/pkg/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//Init router
func Init() {
	c := config.GetConfig()
	//FWはechoを使用
	e := echo.New()

	// CORSの設定追加。下記のような形で設定
	// AllowOrigins: []string{"https://labstack.net"},
	e.Use(middleware.CORS())
	//echoのmiddleware リクエスト単位のログを出力
	e.Use(middleware.Logger())
	//echoのmiddleware 予期せずpanic時、サーバは落とさずにエラーを返せるようにリカバリーする
	e.Use(middleware.Recover())

	//DB接続テスト用
	e.GET("/", clinic_api.Hello)
	e.GET("/users", user_api.ShowUsers)
	// Routing
	e.POST("/signup", handler.Signup)
	e.POST("/login", handler.Login)

	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.POST("", handler.Restricted())

	//ルーティングをグループ化できる
	v1 := e.Group("/api/v1")
	{
		//グループ下のルートではでJWTによる認証必須とする。
		v1.Use(middleware.JWT([]byte("secret")))
		v1.POST("/current-user", user_api.GetCurrentUser)
		v1.POST("/user/:user_id", user_api.ShowUser)
		v1.POST("/users", user_api.ShowUsers)
	}
	v2 := e.Group("/api/v2")
	{
		v2.Use(middleware.JWT([]byte("secret")))
	}

	//e.start(ポート番号)でサーバースタート
	e.Logger.Fatal(e.Start(":" + c.GetString("server.port")))
}
