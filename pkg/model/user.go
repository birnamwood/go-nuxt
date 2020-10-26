package model

import (
	"time"

	"github.com/birnamwood/go-nuxt/pkg/database"
	"gorm.io/gorm"
)

// User Struct gorm.modelと書くと、IDと~_at系のフィールドができる
type User struct {
	gorm.Model
	Email    string
	Password string
	Name     string
	Token    string
}

func (u *User) Create() (err error) {
	db := database.GetDB()
	return db.Create(u).Error
}

func (u *User) Update() (err error) {
	db := database.GetDB()
	return db.Save(u).Error
}

func (u *User) Delete() (err error) {
	db := database.GetDB()
	return db.Delete(u).Error
}

func (u *User) DeleteUpdate() (err error) {
	db := database.GetDB()
	return db.UpdateColumn("DeletedAt", time.Now()).Error
}
