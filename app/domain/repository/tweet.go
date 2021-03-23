package repository

import (
	"github.com/1k-ct/twitter-dem/app/domain/model"
	"github.com/gin-gonic/gin"
)

type TweetRepository interface {
	RegisterTweet(c *gin.Context, tweet *model.Tweet) error
	GetTweetByID(c *gin.Context, id int64) (*model.Tweet, error)
	GetTweetByIDLimitOffset(c *gin.Context, id, limit, offset int64) ([]*model.Tweet, error)
	DeleteTweetByID(c *gin.Context, id int64) (*model.Response, error)
}
