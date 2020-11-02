package route

import (
	"net"
	"os"

	"github.com/birnamwood/go-nuxt/config"
	"github.com/birnamwood/go-nuxt/initialize/database"
	"github.com/birnamwood/go-nuxt/initialize/di"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//Init router
func Init() {
	c := config.GetConfig()

	//依存性の注入
	userHandler := di.InitializeUserHandler(database.GetDB())

	//FWはechoを使用
	e := echo.New()
	//起動時にログにバナーを表示しない
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
		Format: httpLogFormat(),
		Output: file,
	}))
	e.Logger.SetOutput(file)

	e.Use(middleware.BodyDump(bodyDumpHandler))
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

	//e.start(ポート番号)でサーバースタート
	e.Logger.Fatal(e.Start(":" + c.GetString("server.port")))
}

func httpLogFormat() string {
	// Refer to https://github.com/tkuchiki/alp
	var format string
	format += "time:${time_rfc3339}\t"
	format += "host:${remote_ip}\t"
	format += "forwardedfor:${header:x-forwarded-for}\t"
	format += "req:-\t"
	format += "status:${status}\t"
	format += "method:${method}\t"
	format += "uri:${uri}\t"
	format += "size:${bytes_out}\t"
	format += "referer:${referer}\t"
	format += "ua:${user_agent}\t"
	format += "reqtime_ns:${latency}\t"
	format += "cache:-\t"
	format += "runtime:-\t"
	format += "apptime:-\t"
	format += "vhost:${host}\t"
	format += "reqtime_human:${latency_human}\t"
	format += "x-request-id:${id}\t"
	format += "host:${host}\n"
	return format
}

func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	c.Echo().Logger.Print("IPアドレス:", net.ParseIP(c.RealIP()), " Request Body: %v\n", string(reqBody))
	c.Echo().Logger.Print("IPアドレス:", net.ParseIP(c.RealIP()), " Response Body: %v\n", string(resBody))
}
