package model

import (
	"github.com/birnamwood/go-nuxt/database"

	"gorm.io/gorm"
)

// User Struct
type User struct {
	gorm.Model
	Email    string
	Password string
	Name     string
	Token    string
}

// Create comment
func (u *User) Create() (err error) {
	db := database.GetDB()
	return db.Create(u).Error
}

// Update comment
func (u *User) Update() (err error) {
	db := database.GetDB()
	return db.Save(u).Error
}

// Delete comment
func (u *User) Delete() (err error) {
	db := database.GetDB()
	return db.Delete(u).Error
}
