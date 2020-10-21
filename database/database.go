package database

//
import (
	"github.com/birnamwood/go-nuxt/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

//Init database接続
func Init(models ...interface{}) {
	//configの内容取得
	c := config.GetConfig()
	var err error
	//configからデータベースのプロバイダとパスを取得しOpenする
	db, err = gorm.Open(c.GetString("db.provider"), c.GetString("db.url"))
	if err != nil {
		panic("failed to connect database")
	}
	//parameterでモデル名が渡された場合マイグレーション
	db.AutoMigrate(models...)
}

//GetDB return db connection
func GetDB() *gorm.DB {
	return db
}

//Close close database
func Close() {
	db.Close()
}
