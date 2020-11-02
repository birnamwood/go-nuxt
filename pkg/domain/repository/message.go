package repository

import (
	"github.com/birnamwood/go-nuxt/pkg/domain/model"
)

//MessageRepository interface
type MessageRepository interface {
	Create(message *model.Message) (*model.Message, error)
	Update(message *model.Message) (*model.Message, error)
	FindAll() ([]*model.Message, error)
	FindByUserID(userID int) (*model.Message, error)
}
