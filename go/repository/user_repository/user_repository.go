package user_repository

import (
	"github.com/birnamwood/go-nuxt/database"
	"github.com/birnamwood/go-nuxt/model"
)

//FindAllUsers レコード全件取得
func FindAllUsers() []model.User {
	db := database.GetDB()
	var users []model.User
	db.Find(&users)
	return users
}

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
