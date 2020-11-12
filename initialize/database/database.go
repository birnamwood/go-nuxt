package database

//
import (
	"github.com/birnamwood/go-nuxt/initialize/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

//Init database接続
func Init() {
	//configの内容取得
	c := config.GetConfig()
	var err error
	dsn := "host=" + c.GetString("db.host") +
		" port=" + c.GetString("db.port") +
		" dbname=" + c.GetString("db.dbname") +
		" user=" + c.GetString("db.user") +
		" password=" + c.GetString("db.password")

	//configからデータベースのプロバイダとパスを取得しOpenする
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("データベースへの接続失敗")
	}
}

//GetDB return db connection
func GetDB() *gorm.DB {
	return db
}
