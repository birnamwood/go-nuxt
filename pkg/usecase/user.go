package usecase

import (
	"github.com/birnamwood/go-nuxt/pkg/domain/model"
	"github.com/birnamwood/go-nuxt/pkg/domain/repository"
)

type UserUsecase interface {
	Create(user *model.User) (*model.User, error)
	// Update(user *model.User) (*model.User, error)
	// DeleteUpdate(user *model.User) (*model.User, error)
	// Delete(user *model.User) error
	FindAll() ([]*model.User, error)
	FindByID(id int) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (uu *userUsecase) Create(user *model.User) (*model.User, error) {
	newUser, err := uu.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
func (uu *userUsecase) FindByEmail(email string) (*model.User, error) {
	user, err := uu.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu *userUsecase) FindByID(id int) (*model.User, error) {
	user, err := uu.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu *userUsecase) FindAll() ([]*model.User, error) {
	user, err := uu.userRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return user, nil
}
