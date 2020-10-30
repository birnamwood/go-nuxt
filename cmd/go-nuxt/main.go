package main

import (
	"flag"

	"github.com/birnamwood/go-nuxt/config"
	"github.com/birnamwood/go-nuxt/initialize/database"
	"github.com/birnamwood/go-nuxt/initialize/logger"
	"github.com/birnamwood/go-nuxt/initialize/route"
	"go.uber.org/zap"
)

// main
func main() {
	// 環境設定取得 flag.String(<パラメータ名>, <デフォルト値>, <パラメータの説明>)
	env := flag.String("e", "development", "動作環境名")
	//変数宣言のあとに、flag.Parseを実行することでコマンドラインのパラメータがパースされ、各変数に値が格納されます
	flag.Parse()
	//パラメータを渡してconfigの初期化を行う
	config.Init(*env)

	logger := logger.Init()
	zap.ReplaceGlobals(logger)
	logger.Info("Logger Initialize")

	database.Init()
	logger.Info("DB Initialize")

	logger.Info("============= Server Srart =============")
	route.Init()
}
