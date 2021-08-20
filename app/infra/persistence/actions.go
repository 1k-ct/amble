package persistence

import (
	"github.com/1k-ct/amble/app/domain/model"
	"github.com/1k-ct/amble/app/domain/repository"
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
func (ap *actionsPersistence) GetLikeUser(toTweetID string) ([]*model.LikedUser, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	likedUsers := []*model.LikedUser{}
	if err := db.Where("to_tweet_id = ?", toTweetID).
		Find(&likedUsers).Error; err != nil {
		return nil, err
	}
	return likedUsers, nil
}
func (ap *actionsPersistence) GetRetweetUser(toTweetID string) ([]*model.RetweetedUser, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	retweetUsers := []*model.RetweetedUser{}
	if err := db.Where("to_tweet_id = ?", toTweetID).
		Find(&retweetUsers).Error; err != nil {
		return nil, err
	}
	return retweetUsers, nil
}
func (ap *actionsPersistence) GetReply(toTweetID string) ([]*model.Reply, error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	replies := []*model.Reply{}
	if err := db.Where("to_tweet_id = ?", toTweetID).
		Find(&replies).Error; err != nil {
		return nil, err
	}
	return replies, nil
}
