package usecase

import (
	"github.com/1k-ct/amble/app/domain/model"
	"github.com/1k-ct/amble/app/domain/repository"
)

type AccountUseCase interface {
	FindByID(ID string) (*model.User, error)
	RegisterUserAccount(user *model.User) error
	GetUserName(staticID string) (string, error)
	EditUserProfile(*model.User) (*model.User, error)
}
type accountUseCase struct {
	accountRepository repository.AccountRepository
}

func NewAccountUseCase(ar repository.AccountRepository) AccountUseCase {
	return &accountUseCase{
		accountRepository: ar,
	}
}
func (au *accountUseCase) FindByID(ID string) (*model.User, error) {
	user, err := au.accountRepository.FindByID(ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (au *accountUseCase) RegisterUserAccount(user *model.User) error {
	if err := au.accountRepository.RegisterUserAccount(user); err != nil {
		return err
	}
	return nil
}
func (au *accountUseCase) GetUserName(staticID string) (string, error) {
	userName, err := au.accountRepository.GetUserName(staticID)
	if err != nil {
		return "", err
	}
	return userName, nil
}
func (au *accountUseCase) EditUserProfile(user *model.User) (*model.User, error) {
	userProfile, err := au.accountRepository.EditUserProfile(user)
	if err != nil {
		return nil, err
	}
	return userProfile, nil
}
