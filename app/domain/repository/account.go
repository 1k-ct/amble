package repository

import "github.com/1k-ct/amble/app/domain/model"

type AccountRepository interface {
	FindByID(ID string) (*model.User, error)
	RegisterUserAccount(*model.User) error
	GetUserName(staticID string) (string, error)
	EditUserProfile(*model.User) (*model.User, error)
}
