package model

import "gorm.io/gorm"

// Card Struct gorm.modelと書くと、IDと~_at系のフィールドができる
type Card struct {
	gorm.Model
	AccountID int `gorm:"foreignKey:AccountRefer,constraint:OnDelete:CASCADE;"`
	Account   Account
	Name      string
}
