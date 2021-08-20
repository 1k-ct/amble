package persistence

import (
	"errors"

	"github.com/1k-ct/amble/app/domain/model"
	"github.com/1k-ct/amble/app/domain/repository"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
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
		// どちら場合でもいい。
		// ここの場合は、この先でエラー処理をする。
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
	user.ID = uuid.NewV4().String()
	if err := db.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}
	return nil
}
func (ap *accountPersistence) GetUserName(staticID string) (string, error) {
	db, err := Connect()
	if err != nil {
		return "", err
	}
	user := &model.User{}
	if err := db.Where("id = ?", staticID).
		Find(&user).Error; err != nil {
		return "", err
	}
	return user.UserName, nil
}
func (ap *accountPersistence) EditUserProfile(userProfile *model.User) (*model.User, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Where("id = ?", userProfile.ID).
		Find(&userProfile).Error; err != nil {
		return nil, err
	}
	user := &model.User{
		UserName:  userProfile.UserName,
		Location:  userProfile.Location,
		FreeSpace: userProfile.FreeSpace,
	}
	if err := db.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
