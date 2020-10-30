package model

import (
	"time"

	"gorm.io/gorm"
)

// User Struct gorm.modelと書くと、IDと~_at系のフィールドができる
type User struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Email     string
	Password  string
	Name      string
	Token     string
}
