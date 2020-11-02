package model

import (
	"gorm.io/gorm"
)

// Device Struct gorm.modelと書くと、IDと~_at系のフィールドができる
type Device struct {
	gorm.Model
	UserID int
	User   User `gorm:"foreignKey:UserRefer,constraint:OnDelete:CASCADE;"`
	Name   string
}
