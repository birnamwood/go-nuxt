package repository

import (
	"../database"
	"../model"
)

//FindUserByID comment
func FindUserByID(id int) model.User {
	db := database.GetDB()
	var user model.User
	db.First(&user, "id = ?", id)
	return user
}

//FindUserByEmail comment
func FindUserByEmail(email string) model.User {
	db := database.GetDB()
	var user model.User
	db.First(&user, "Email = ?", email)
	return user
}
