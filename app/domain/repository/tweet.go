package repository

import (
	"github.com/1k-ct/twitter-dem/app/domain/model"
)

type TweetRepository interface {
	RegisterTweet(tweet *model.Tweet) error
	GetTweetByID(staticID string) (*model.Tweet, error)
	GetTweetByIDs(ids []int64) ([]*model.Tweet, error)
	DeleteTweetByID(id int64) error
}
