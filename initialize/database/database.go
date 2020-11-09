package database

//
import (
	"log"

	"github.com/birnamwood/go-nuxt/initialize/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
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
	m, err := migrate.New(
		"file://migration",
		"postgres://postgres:postgres@postgres:5432/db?sslmode=disable")
	if err != nil {
		log.Fatal("マイグレーション失敗")
	}
	steperr := m.Steps(2)
	if steperr != nil {
		zap.S().Info(steperr.Error())
	}
}

//GetDB return db connection
func GetDB() *gorm.DB {
	return db
}
