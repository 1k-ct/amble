package usecase

import (
	"github.com/1k-ct/amble/app/domain/model"
	"github.com/1k-ct/amble/app/domain/repository"
)

type ActionsUseCase interface {
	Like(userStaticID, staticID string) error
	Retweet(userStaticID, staticID string) error
	Reply(reply *model.Reply) error
	GetLikeUser(toTweetID string) ([]*model.LikedUser, error)
	GetRetweetUser(toTweetID string) ([]*model.RetweetedUser, error)
	GetReply(toTweetID string) ([]*model.Reply, error)
}
type actionsUseCase struct {
	actionsUseCase repository.ActionsRepository
}

func NewActionsUseCase(ar repository.ActionsRepository) ActionsUseCase {
	return &actionsUseCase{
		actionsUseCase: ar,
	}
}

func (au *actionsUseCase) Like(userStaticID, staticID string) error {
	if err := au.actionsUseCase.Like(userStaticID, staticID); err != nil {
		return err
	}
	return nil
}
func (au *actionsUseCase) Retweet(userStaticID, staticID string) error {
	if err := au.actionsUseCase.Retweet(userStaticID, staticID); err != nil {
		return err
	}
	return nil
}
func (au *actionsUseCase) Reply(reply *model.Reply) error {
	if err := au.actionsUseCase.Reply(reply); err != nil {
		return err
	}
	return nil
}
func (au *actionsUseCase) GetLikeUser(toTweetID string) ([]*model.LikedUser, error) {
	likedUsers, err := au.actionsUseCase.GetLikeUser(toTweetID)
	if err != nil {
		return nil, err
	}
	return likedUsers, nil
}
func (au *actionsUseCase) GetRetweetUser(toTweetID string) ([]*model.RetweetedUser, error) {
	retweetUsers, err := au.actionsUseCase.GetRetweetUser(toTweetID)
	if err != nil {
		return nil, err
	}
	return retweetUsers, nil
}
func (au *actionsUseCase) GetReply(toTweetID string) ([]*model.Reply, error) {
	replies, err := au.actionsUseCase.GetReply(toTweetID)
	if err != nil {
		return nil, err
	}
	return replies, nil
}
