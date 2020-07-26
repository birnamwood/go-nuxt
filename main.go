package main

import (
	"flag"

	"./config"
	"./database"
	"./route"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// main
func main() {
	// 環境設定取得
	env := flag.String("e", "development", "")
	flag.Parse()

	config.Init(*env)
	database.Init()
	route.Init()
}
