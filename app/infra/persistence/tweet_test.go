package persistence

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/1k-ct/twitter-dem/app/domain/model"
	"github.com/1k-ct/twitter-dem/pkg/database"
	"github.com/gin-gonic/gin"
)

func TestRegisterTweet(t *testing.T) {
	db, err := database.NewLocalDB("user", "password", "sample")
	if err != nil {
		t.Error(err)
	}
	if err := db.NewMakeDB(&model.Tweet{}); err != nil {
		t.Error(err)
	}

	// db, err := connect()
	// if err != nil {
	// 	t.Error(err)
	// }
	// twe := &model.Tweet{}
	// db.AutoMigrate(twe)

	tweet := &model.Tweet{
		// ID:           0,
		Name:         "sato",
		Content:      "テストのでーたツイート内容",
		LikeCount:    0,
		RetweetCount: 0,
		ReplyCount:   0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	var c *gin.Context
	if err := NewTweetPersistence().RegisterTweet(c, tweet); err != nil {
		t.Error(err)
	}
}

func TestGetTweetByID(t *testing.T) {
	var c *gin.Context
	tweet, err := NewTweetPersistence().GetTweetByID(c, 2)
	if err != nil {
		t.Error(err)
	}
	log.Println(tweet)
}

func TestGetTweetByIDLimitOffset(t *testing.T) {
	var c *gin.Context
	tweets, err := NewTweetPersistence().GetTweetByIDs(c, []int64{1, 2})
	if err != nil {
		t.Error(err)
	}
	for _, tweet := range tweets {
		fmt.Println(tweet)
	}
	// log.Println(tweets)
}
func TestDeleteTweetByID(t *testing.T) {
	var c *gin.Context
	err := NewTweetPersistence().DeleteTweetByID(c, 1)
	if err != nil {
		t.Error(err)
	}
}
func TestGetTweetByIDs(t *testing.T) {
	var c *gin.Context
	tweets, err := NewTweetPersistence().GetTweetByIDs(c, []int64{1, 2})
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < len(tweets); i++ {
		if tweets[i].ID == int64(i) {
			t.Error(err)
		}
	}
}
