package persistence

import (
	"time"

	"github.com/birnamwood/go-nuxt/pkg/domain/model"
	"github.com/birnamwood/go-nuxt/pkg/domain/repository"
	"gorm.io/gorm"
)

type UserPersistence struct {
	db *gorm.DB
}

func NewUserPersistence(database *gorm.DB) repository.UserRepository {
	return &UserPersistence{db: database}
}

func (up *UserPersistence) Create(user *model.User) (*model.User, error) {
	if err := up.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (up *UserPersistence) Update(user *model.User) (*model.User, error) {
	if err := up.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (up *UserPersistence) Delete(user *model.User) error {
	return up.db.Delete(&user).Error
}

func (up *UserPersistence) DeleteUpdate(user *model.User) (*model.User, error) {
	if err := up.db.UpdateColumn("DeletedAt", time.Now()).Error; err != nil {
		return nil, err
	}
	return user, nil
}

//FindAll レコード全件取得
func (up *UserPersistence) FindAll() ([]*model.User, error) {
	users := []*model.User{}
	if err := up.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

//FindByID comment
func (up *UserPersistence) FindByID(id int) (*model.User, error) {
	var user *model.User
	if err := up.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

//FindByEmail comment
func (up *UserPersistence) FindByEmail(email string) (*model.User, error) {
	var user *model.User
	if err := up.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
