package repository

import "github.com/1k-ct/twitter-dem/app/domain/model"

type ActionsRepository interface {
	Like(userStaticID, staticID string) error
	Retweet(userStaticID, staticID string) error
	Reply(reply *model.Reply) error
}
