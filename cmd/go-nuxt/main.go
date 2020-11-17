package main

import (
	"flag"

	"github.com/birnamwood/go-nuxt/initialize/config"
	"github.com/birnamwood/go-nuxt/initialize/database"
	"github.com/birnamwood/go-nuxt/initialize/firebase"
	"github.com/birnamwood/go-nuxt/initialize/logger"
	"github.com/birnamwood/go-nuxt/initialize/route"
	"go.uber.org/zap"
)

func main() {
	// 環境設定取得 flag.String(<パラメータ名>, <デフォルト値>, <パラメータの説明>)
	env := flag.String("e", "development", "動作環境名")
	//変数宣言のあとに、flag.Parseを実行することでコマンドラインのパラメータがパースされ、各変数に値が格納されます
	flag.Parse()
	//パラメータを渡してconfigの初期化を行う
	config.Init(*env, "./initialize/config/environments/")

	//Loggerの初期化
	logger := logger.Init()
	zap.ReplaceGlobals(logger)
	logger.Info("Logger Initialize")

	//Database初期化
	database.Init()
	logger.Info("DB Initialize")

	//Firebase初期化
	firebase.Init()
	logger.Info("Firebase Initialize")

	route.Init()
}
