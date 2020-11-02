// +build wireinject

package di

import (
	"github.com/birnamwood/go-nuxt/pkg/infrastructure/persistence"
	"github.com/birnamwood/go-nuxt/pkg/interface/handler"
	"github.com/birnamwood/go-nuxt/pkg/usecase"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeUserHandler(db *gorm.DB) handler.UserHandler {
	wire.Build(handler.NewUserHandler, usecase.NewUserUsecase, persistence.NewUserPersistence)
	return nil
}
