package model

// User comment
type User struct {
	ID       int    `'json:"id" gorm:"primary_key"`
	Mail     string `json:"mail"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// CreateUser comment
func CreateUser(user *User) {
	db.Create(user)
}

// FindUser comment
func FindUser(u *User) User {
	var user User
	db.Where(u).First(&user)
	return user
}
