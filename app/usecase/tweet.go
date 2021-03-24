package usecase

import (
	"github.com/1k-ct/twitter-dem/app/domain/model"
	"github.com/1k-ct/twitter-dem/app/domain/repository"
	"github.com/gin-gonic/gin"
)

type TweetUseCase interface {
	RegisterTweet(c *gin.Context, tweet *model.Tweet) error
	GetTweetByID(c *gin.Context, id int64) (*model.Tweet, error)
	GetTweetByIDs(c *gin.Context, ids []int64) ([]*model.Tweet, error)
	DeleteTweetByID(c *gin.Context, id int64) error
}
type tweetUseCase struct {
	tweetRepository repository.TweetRepository
}

func NewTweetUseCase(tr repository.TweetRepository) TweetUseCase {
	return &tweetUseCase{
		tweetRepository: tr,
	}
}

func (tu *tweetUseCase) RegisterTweet(c *gin.Context, tweet *model.Tweet) error {
	if err := tu.tweetRepository.RegisterTweet(c, tweet); err != nil {
		return err
	}
	return nil
}
func (tu *tweetUseCase) GetTweetByID(c *gin.Context, id int64) (*model.Tweet, error) {
	tweet, err := tu.tweetRepository.GetTweetByID(c, id)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}
func (tu *tweetUseCase) GetTweetByIDs(c *gin.Context, ids []int64) ([]*model.Tweet, error) {
	tweet, err := tu.tweetRepository.GetTweetByIDs(c, ids)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}
func (tu *tweetUseCase) DeleteTweetByID(c *gin.Context, id int64) error {
	err := tu.tweetRepository.DeleteTweetByID(c, id)
	if err != nil {
		return err
	}
	return nil
}
