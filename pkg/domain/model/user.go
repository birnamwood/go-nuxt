package model

import (
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
