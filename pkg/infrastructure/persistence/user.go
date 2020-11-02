package persistence

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/birnamwood/go-nuxt/pkg/domain/model"
	"github.com/birnamwood/go-nuxt/pkg/domain/repository"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//UserPersistence struct
type UserPersistence struct {
	db *gorm.DB
}

//NewUserPersistence New
func NewUserPersistence(database *gorm.DB) repository.UserRepository {
	return &UserPersistence{db: database}
}

//Create UserCreate
func (up *UserPersistence) Create(user *model.User) (*model.User, error) {
	if err := up.db.Create(&user).Error; err != nil {
		json, _ := json.Marshal(&user)
		return nil, errors.Wrapf(err, "ユーザー情報の更新に失敗しました。 User: '%s'", string(json))

	}
	return user, nil
}

//Update Updateuser
func (up *UserPersistence) Update(user *model.User) (*model.User, error) {
	if err := up.db.Save(&user).Error; err != nil {
		json, _ := json.Marshal(&user)
		return nil, errors.Wrapf(err, "ユーザー情報の更新に失敗しました。 User: '%s'", string(json))
	}
	return user, nil
}

//Delete User
func (up *UserPersistence) Delete(user *model.User) error {
	if err := up.db.Delete(&user).Error; err != nil {
		json, _ := json.Marshal(&user)
		return errors.Wrapf(err, "ユーザー情報の削除に失敗しました。 User: '%s'", string(json))
	}
	return nil
}

//DeleteUpdate User論理削除
func (up *UserPersistence) DeleteUpdate(user *model.User) (*model.User, error) {
	if err := up.db.Model(&user).Update("DeletedAt", time.Now()).Error; err != nil {
		json, _ := json.Marshal(&user)
		return nil, errors.Wrapf(err, "ユーザーの論理削除に失敗しました。User: '%s'", string(json))
	}
	return user, nil
}

//FindAll レコード全件取得
func (up *UserPersistence) FindAll() ([]*model.User, error) {
	users := []*model.User{}
	if err := up.db.Find(&users).Error; err != nil {
		return nil, errors.Wrap(err, "全ユーザー検索に失敗しました。")
	}
	return users, nil
}

//FindByID comment
func (up *UserPersistence) FindByID(id int) (*model.User, error) {
	user := &model.User{ID: id}
	if err := up.db.Where("ID = ?", id).First(&user).Error; err != nil {
		return nil, errors.Wrapf(err, "IDでのユーザー検索に失敗しました。: '%s'", strconv.Itoa(id))
	}
	return user, nil
}

//FindByEmail comment
func (up *UserPersistence) FindByEmail(email string) (*model.User, error) {
	user := &model.User{Email: email}
	if err := up.db.Where("email = ?", email).First(&user).Error; err != nil {
		zap.S().Warn("err: ", "Emailでのユーザー検索に失敗しました。:", email)
		return nil, errors.Wrapf(err, "Emailでのユーザー検索に失敗しました。: '%s'", email)
	}
	return user, nil
}
