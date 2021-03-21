package persistence

import (
	"github.com/1k-ct/twitter-dem/app/domain/model"
	"github.com/1k-ct/twitter-dem/app/domain/repository"
	"github.com/gin-gonic/gin"
)

type tweetPersistence struct{}

func NewTweetPersistence() repository.TweetRepository {
	return &tweetPersistence{}
}

func (tp *tweetPersistence) RegisterTweet(c *gin.Context) (*model.Response, error) {
	return nil, nil
}
func (tp *tweetPersistence) GetTweetByIDs(c *gin.Context, id []int64) ([]*model.Tweet, error) {
	return nil, nil
}
func (tp *tweetPersistence) DeleteTweetByID(c *gin.Context, id int64) (*model.Response, error) {
	return nil, nil
}
