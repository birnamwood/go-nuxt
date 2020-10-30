package route

import (
	"os"

	"github.com/birnamwood/go-nuxt/config"
	"github.com/birnamwood/go-nuxt/initialize/database"
	"github.com/birnamwood/go-nuxt/pkg/infrastructure/persistence"
	"github.com/birnamwood/go-nuxt/pkg/interface/handler"
	"github.com/birnamwood/go-nuxt/pkg/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//Init router
func Init() {
	c := config.GetConfig()

	//依存性の注入
	userPersistence := persistence.NewUserPersistence(database.GetDB())
	userUsecase := usecase.NewUserUsecase(userPersistence)
	userHandler := handler.NewUserHandler(userUsecase)

	//FWはechoを使用
	e := echo.New()
	e.HideBanner = true

	// CORSの設定追加。下記のような形で設定
	// AllowOrigins: []string{"https://labstack.net"},
	e.Use(middleware.CORS())
	//echoのmiddleware リクエスト単位のログを出力
	file, err := os.OpenFile(c.GetString("log.filename2"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: file,
	}))
	e.Logger.SetOutput(file)
	//echoのmiddleware 予期せずpanic時、サーバは落とさずにエラーを返せるようにリカバリーする
	e.Use(middleware.Recover())

	//DB接続テスト用
	e.GET("/", userHandler.HelloUser)
	e.GET("/users", userHandler.ShowUsers)
	e.GET("/user/:user_id", userHandler.ShowUser)
	// Routing
	e.POST("/signup", userHandler.Signup)
	e.POST("/login", userHandler.Login)

	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.POST("", userHandler.Restricted)

	//ルーティングをグループ化できる
	v1 := e.Group("/api/v1")
	{
		//グループ下のルートではでJWTによる認証必須とする。
		v1.Use(middleware.JWT([]byte("secret")))
		v1.POST("/current-user", userHandler.GetCurrentUser)
		v1.POST("/user/:user_id", userHandler.ShowUser)
		v1.POST("/users", userHandler.ShowUsers)
	}
	v2 := e.Group("/api/v2")
	{
		v2.Use(middleware.JWT([]byte("secret")))
	}

	e.Logger.Info("Server Start" + ":" + c.GetString("server.port"))
	//e.start(ポート番号)でサーバースタート
	e.Logger.Fatal(e.Start(":" + c.GetString("server.port")))
}
