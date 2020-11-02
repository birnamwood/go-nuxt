package repository

import (
	"github.com/birnamwood/go-nuxt/pkg/domain/model"
)

//UserRepository text
type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	DeleteUpdate(user *model.User) (*model.User, error)
	Delete(user *model.User) error
	FindAll() ([]*model.User, error)
	FindByID(id int) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}
