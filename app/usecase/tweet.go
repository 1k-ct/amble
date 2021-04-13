package usecase

import (
	"github.com/1k-ct/twitter-dem/app/domain/model"
	"github.com/1k-ct/twitter-dem/app/domain/repository"
)

type TweetUseCase interface {
	RegisterTweet(tweet *model.Tweet) error
	GetTweetByID(staticID string) (*model.Tweet, error)
	GetTweetByIDs(ids []int64) ([]*model.Tweet, error)
	DeleteTweetByID(id int64) error
}
type tweetUseCase struct {
	tweetRepository repository.TweetRepository
}

func NewTweetUseCase(tr repository.TweetRepository) TweetUseCase {
	return &tweetUseCase{
		tweetRepository: tr,
	}
}

func (tu *tweetUseCase) RegisterTweet(tweet *model.Tweet) error {
	if err := tu.tweetRepository.RegisterTweet(tweet); err != nil {
		return err
	}
	return nil
}
func (tu *tweetUseCase) GetTweetByID(staticID string) (*model.Tweet, error) {
	tweet, err := tu.tweetRepository.GetTweetByID(staticID)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}
func (tu *tweetUseCase) GetTweetByIDs(ids []int64) ([]*model.Tweet, error) {
	tweet, err := tu.tweetRepository.GetTweetByIDs(ids)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}
func (tu *tweetUseCase) DeleteTweetByID(id int64) error {
	err := tu.tweetRepository.DeleteTweetByID(id)
	if err != nil {
		return err
	}
	return nil
}
