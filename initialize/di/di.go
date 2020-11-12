// +build wireinject

package di

import (
	"github.com/birnamwood/go-nuxt/pkg/infrastructure/persistence"
	"github.com/birnamwood/go-nuxt/pkg/interface/handler"
	"github.com/birnamwood/go-nuxt/pkg/usecase"
	"github.com/google/wire"
	"gorm.io/gorm"
)

//InitializeUserHandler を依存関係から生成
func InitializeUserHandler(db *gorm.DB) handler.UserHandler {
	wire.Build(handler.NewUserHandler, usecase.NewUserUsecase, persistence.NewUserPersistence)
	return nil
}

//InitializeMessageHandler を依存関係から生成
func InitializeMessageHandler(db *gorm.DB) handler.MessageHandler {
	wire.Build(handler.NewMessageHandler, usecase.NewMessageUsecase, persistence.NewMessagePersistence)
	return nil
}
