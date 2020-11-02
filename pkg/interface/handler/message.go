package handler

import (
	"net/http"

	"github.com/birnamwood/go-nuxt/pkg/domain/model"
	"github.com/birnamwood/go-nuxt/pkg/usecase"
	"github.com/labstack/echo/v4"
)

//MessageHandler interface
type MessageHandler interface {
	SendMessage(echo.Context) error
}

//messageHandler struct
type messageHandler struct {
	messageUsecase usecase.MessageUsecase
}

//NewMessageHandler New
func NewMessageHandler(messageUsecase usecase.MessageUsecase) MessageHandler {
	return &messageHandler{messageUsecase: messageUsecase}
}

//SendMessage
func (mh *messageHandler) SendMessage(c echo.Context) error {
	message := new(model.Message)
	if err := c.Bind(message); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	createMessage, createErr := mh.messageUsecase.Create(message)
	if createErr != nil {
		return c.JSON(http.StatusBadRequest, createErr.Error())
	}
	return c.JSON(http.StatusOK, createMessage)
}
