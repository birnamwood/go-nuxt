package database

//
import (
	"../config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

//Init database接続
func Init(models ...interface{}) {
	c := config.GetConfig()
	var err error
	db, err = gorm.Open(c.GetString("db.provider"), c.GetString("db.url"))
	if err != nil {
		panic("failed to connect database")
	}
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
