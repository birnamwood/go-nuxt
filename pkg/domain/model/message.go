package model

import (
	"time"

	"gorm.io/gorm"
)

// Message Struct gorm.modelと書くと、IDと~_at系のフィールドができる
type Message struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    int
	User      User `gorm:"foreignKey:UserRefer,constraint:OnDelete:CASCADE;"`
	Title     string
	Body      string
}
