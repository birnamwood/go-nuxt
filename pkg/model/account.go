package model

import (
	"time"

	"github.com/birnamwood/go-nuxt/pkg/database"
	"gorm.io/gorm"
)

// Account Struct gorm.modelと書くと、IDと~_at系のフィールドができる
type Account struct {
	gorm.Model
	UserID int
	User   User `gorm:"foreignKey:UserRefer,constraint:OnDelete:CASCADE;"`
	Name   string
}

func (a *Account) Create() (err error) {
	db := database.GetDB()
	return db.Create(a).Error
}

func (a *Account) Update() (err error) {
	db := database.GetDB()
	return db.Save(a).Error
}

func (a *Account) Delete() (err error) {
	db := database.GetDB()
	return db.Delete(a).Error
}

func (a *Account) DeleteUpdate() (err error) {
	db := database.GetDB()
	return db.UpdateColumn("DeletedAt", time.Now()).Error
}
