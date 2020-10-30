package model

import (
	"time"

	"github.com/birnamwood/go-nuxt/initialize/database"
	"gorm.io/gorm"
)

// Account Struct gorm.modelと書くと、IDと~_at系のフィールドができる
type Device struct {
	gorm.Model
	UserID int
	User   User `gorm:"foreignKey:UserRefer,constraint:OnDelete:CASCADE;"`
	Name   string
}

func (d *Device) Create() (err error) {
	db := database.GetDB()
	return db.Create(d).Error
}

func (d *Device) Update() (err error) {
	db := database.GetDB()
	return db.Save(d).Error
}

func (d *Device) Delete() (err error) {
	db := database.GetDB()
	return db.Delete(d).Error
}

func (d *Device) DeleteUpdate() (err error) {
	db := database.GetDB()
	return db.UpdateColumn("DeletedAt", time.Now()).Error
}
