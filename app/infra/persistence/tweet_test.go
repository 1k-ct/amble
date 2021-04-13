package persistence

import (
	"errors"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/1k-ct/twitter-dem/app/domain/model"
	"github.com/1k-ct/twitter-dem/pkg/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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
		Name:      "sato",
		Content:   "テストのでーたツイート内容",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := NewTweetPersistence().RegisterTweet(tweet); err != nil {
		t.Error(err)
	}
}

func TestGetTweetByID(t *testing.T) {
	tweet, err := NewTweetPersistence().GetTweetByID("")
	if err != nil {
		t.Error(err)
	}
	log.Println(tweet)
}

func Test_tweetPersistence_GetTweetByID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name      string
		args      args
		isPrivate bool
		wantErr   bool
	}{
		{
			name: "GetTweetID正常",
			args: args{
				id: "6",
			},
			isPrivate: false,
			wantErr:   false,
		},
		{
			name: "GetTweetID異常",
			args: args{
				id: "1",
			},
			isPrivate: true,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTweetPersistence().GetTweetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("tweetPersistence.GetTweetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				if errors.Is(err, gorm.ErrRecordNotFound) != tt.wantErr {
					t.Errorf("tweetPersistence.GetTweetByID() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if !reflect.DeepEqual(got.IsPrivate, tt.isPrivate) {
					t.Errorf("tweetPersistence.GetTweetByID() = %v, want %v", got, tt.isPrivate)
				}
			}
		})
	}
}

// func TestGetTweetByIDLimitOffset(t *testing.T) {
// 	var c *gin.Context
// 	tweets, err := NewTweetPersistence().GetTweetByIDs(c, []int64{1, 2})
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	for _, tweet := range tweets {
// 		fmt.Println(tweet)
// 	}
// 	// log.Println(tweets)
// }
func TestDeleteTweetByID(t *testing.T) {
	err := NewTweetPersistence().DeleteTweetByID(1)
	if err != nil {
		t.Error(err)
	}
}
func TestGetTweetByIDs(t *testing.T) {
	tweets, err := NewTweetPersistence().GetTweetByIDs([]int64{1, 2, 3, 4, 5})
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < len(tweets); i++ {
		if tweets[i].ID == int64(i) {
			t.Error(err)
		}
	}
	for _, tweet := range tweets {
		log.Println(tweet)
	}
}
