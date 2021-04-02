package persistence

import (
	"errors"

	"github.com/1k-ct/twitter-dem/app/domain/model"
	"github.com/1k-ct/twitter-dem/app/domain/repository"
	"github.com/jinzhu/gorm"
)

type accountPersistence struct{}

func NewAccountPersistence() repository.AccountRepository {
	return &accountPersistence{}
}

func (ap *accountPersistence) FindByID(ID string) (*model.User, error) {
	user := &model.User{}
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	if err := db.Where("id = ?", ID).Find(&user).Error; err != nil {
		// if errors.Is(err, gorm.ErrRecordNotFound) {
		// 	return nil, err
		// }
		// どちらか
		// if gorm.IsRecordNotFoundError(err) {
		// 	log.Println(err.Error())
		// 	return nil, err
		// }
		return nil, err
	}
	return user, nil
}
func (ap *accountPersistence) RegisterUserAccount(user *model.User) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	if err := db.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}
	return nil
}
