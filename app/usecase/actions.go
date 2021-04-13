package usecase

import (
	"github.com/1k-ct/twitter-dem/app/domain/model"
	"github.com/1k-ct/twitter-dem/app/domain/repository"
)

type ActionsUseCase interface {
	Like(userStaticID, staticID string) error
	Retweet(userStaticID, staticID string) error
	Reply(reply *model.Reply) error
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
