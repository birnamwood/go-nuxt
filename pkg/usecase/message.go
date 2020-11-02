package usecase

import (
	"github.com/birnamwood/go-nuxt/pkg/domain/model"
	"github.com/birnamwood/go-nuxt/pkg/domain/repository"
)

//MessageUsecase interface
type MessageUsecase interface {
	Create(message *model.Message) (*model.Message, error)
	// Update(Message *model.Message) (*model.Message, error)
	FindAll() ([]*model.Message, error)
	FindByUserID(userID int) (*model.Message, error)
}

//Messageusecase struct
type messageUsecase struct {
	messageRepo repository.MessageRepository
}

//NewMessageUsecase New
func NewMessageUsecase(messageRepo repository.MessageRepository) MessageUsecase {
	return &messageUsecase{messageRepo: messageRepo}
}

//Create MessageCreate
func (mu *messageUsecase) Create(Message *model.Message) (*model.Message, error) {
	newMessage, err := mu.messageRepo.Create(Message)
	if err != nil {
		return nil, err
	}
	return newMessage, nil
}

//FindAll FindMessages
func (mu *messageUsecase) FindAll() ([]*model.Message, error) {
	Message, err := mu.messageRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return Message, nil
}

//FindByUserId FindMessage
func (mu *messageUsecase) FindByUserID(userID int) (*model.Message, error) {
	Message, err := mu.messageRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	return Message, nil
}
