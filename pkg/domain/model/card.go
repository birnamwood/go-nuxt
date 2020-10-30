package model

import (
	"time"

	"github.com/birnamwood/go-nuxt/initialize/database"
	"gorm.io/gorm"
)

// Account Struct gorm.modelと書くと、IDと~_at系のフィールドができる
type Card struct {
	gorm.Model
	AccountId int `gorm:"foreignKey:AccountRefer,constraint:OnDelete:CASCADE;"`
	Account   Account
	Name      string
}

func (c *Card) Create() (err error) {
	db := database.GetDB()
	return db.Create(c).Error
}

func (c *Card) Update() (err error) {
	db := database.GetDB()
	return db.Save(c).Error
}

func (c *Card) Delete() (err error) {
	db := database.GetDB()
	return db.Delete(c).Error
}

func (c *Card) DeleteUpdate() (err error) {
	db := database.GetDB()
	return db.UpdateColumn("DeletedAt", time.Now()).Error
}
