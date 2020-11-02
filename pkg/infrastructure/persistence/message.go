package persistence

import (
	"encoding/json"
	"strconv"

	"github.com/birnamwood/go-nuxt/pkg/domain/model"
	"github.com/birnamwood/go-nuxt/pkg/domain/repository"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

//MessagePersistence struct
type MessagePersistence struct {
	db *gorm.DB
}

//NewMessagePersistence new
func NewMessagePersistence(database *gorm.DB) repository.MessageRepository {
	return &MessagePersistence{db: database}
}

//Create Message Create
func (mp *MessagePersistence) Create(message *model.Message) (*model.Message, error) {
	if err := mp.db.Create(&message).Error; err != nil {
		json, _ := json.Marshal(&message)
		return nil, errors.Wrapf(err, "メッセージの送信に失敗しました。 Message: '%s'", string(json))

	}
	return message, nil
}

//Update messageUpdate
func (mp *MessagePersistence) Update(message *model.Message) (*model.Message, error) {
	if err := mp.db.Save(&message).Error; err != nil {
		json, _ := json.Marshal(&message)
		return nil, errors.Wrapf(err, "ユーザー情報の更新に失敗しました。 User: '%s'", string(json))
	}
	return message, nil
}

//FindAll レコード全件取得
func (mp *MessagePersistence) FindAll() ([]*model.Message, error) {
	message := []*model.Message{}
	if err := mp.db.Find(&message).Error; err != nil {
		return nil, errors.Wrap(err, "全ユーザー検索に失敗しました。")
	}
	return message, nil
}

//FindByUserID comment
func (mp *MessagePersistence) FindByUserID(userID int) (*model.Message, error) {
	message := &model.Message{ID: userID}
	if err := mp.db.Where("UserID = ?", userID).First(&message).Error; err != nil {
		return nil, errors.Wrapf(err, "ユーザーIDでの検索に失敗しました。: '%s'", strconv.Itoa(userID))
	}
	return message, nil
}
