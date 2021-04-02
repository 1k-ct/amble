package usecase

import (
	"github.com/1k-ct/twitter-dem/app/domain/model"
	"github.com/1k-ct/twitter-dem/app/domain/repository"
)

type AccountUseCase interface {
	FindByID(ID string) (*model.User, error)
	RegisterUserAccount(user *model.User) error
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
