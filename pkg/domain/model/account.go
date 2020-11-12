package model

import "gorm.io/gorm"

// Account Struct gorm.modelと書くと、IDと~_at系のフィールドができる
type Account struct {
	gorm.Model
	UserID int
	User   User `gorm:"foreignKey:UserRefer,constraint:OnDelete:CASCADE;"`
	Name   string
}
