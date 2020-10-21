package main

import (
	"flag"

	"github.com/birnamwood/go-nuxt/config"
	"github.com/birnamwood/go-nuxt/database"
	"github.com/birnamwood/go-nuxt/route"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// main
func main() {
	// 環境設定取得 flag.String(<パラメータ名>, <デフォルト値>, <パラメータの説明>)
	env := flag.String("e", "development", "動作環境名")
	//変数宣言のあとに、flag.Parseを実行することでコマンドラインのパラメータがパースされ、各変数に値が格納されます
	flag.Parse()
	//パラメータを渡してconfigの初期化を行う
	config.Init(*env)
	database.Init()
	route.Init()
}
