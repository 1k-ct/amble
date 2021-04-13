package persistence

import (
	"github.com/1k-ct/twitter-dem/app/domain/model"
	"github.com/1k-ct/twitter-dem/app/domain/repository"
	uuid "github.com/satori/go.uuid"
)

type actionsPersistence struct{}

func NewActionsPersistence() repository.ActionsRepository {
	return &actionsPersistence{}
}
func (ap *actionsPersistence) Like(staticID, userStaticID string) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	tweet := &model.Tweet{}

	if err := db.Where("static_id = ? AND is_private = ?", staticID, false).
		Find(&tweet).Error; err != nil {
		return err
	}
	tweet.LikedCount += 1
	if err := db.Save(&tweet).Error; err != nil {
		return err
	}
	likedUser := &model.LikedUser{}

	likedUser.ToTweetID = staticID
	likedUser.UserStaticID = userStaticID
	if err := db.Create(&likedUser).Error; err != nil {
		return err
	}
	return nil
}
func (ap *actionsPersistence) Retweet(staticID, userStaticID string) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	tweet := &model.Tweet{}

	if err := db.Where("static_id = ? AND is_private = ?", staticID, false).
		Find(&tweet).Error; err != nil {
		return err
	}
	tweet.RetweetedCount += 1
	if err := db.Save(&tweet).Error; err != nil {
		return err
	}

	retweetedUser := &model.RetweetedUser{}

	retweetedUser.ToTweetID = staticID
	retweetedUser.UserStaticID = userStaticID
	if err := db.Create(&retweetedUser).Error; err != nil {
		return err
	}
	return nil
}
func (ap *actionsPersistence) Reply(reply *model.Reply) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	tweet := &model.Tweet{}

	if err := db.Where("static_id = ? AND is_private = ?", reply.ToTweetID, false).
		Find(&tweet).Error; err != nil {
		return err
	}
	tweet.ReplyCount += 1
	if err := db.Save(&tweet).Error; err != nil {
		return err
	}

	reply.StaticID = uuid.NewV4().String()
	if err := db.Create(&reply).Error; err != nil {
		return err
	}
	return nil
}
