package persistence

import (
	"github.com/1k-ct/twitter-dem/app/domain/model"
	"github.com/1k-ct/twitter-dem/app/domain/repository"
	"github.com/1k-ct/twitter-dem/pkg/database"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type tweetPersistence struct{}

func NewTweetPersistence() repository.TweetRepository {
	return &tweetPersistence{}
}
func connect() (*gorm.DB, error) {
	config, err := database.NewLocalDB("user", "password", "sample")
	if err != nil {
		return nil, err
	}

	db, err := config.Connect()
	if err != nil {
		return nil, err
	}
	return db, nil
}
func (tp *tweetPersistence) RegisterTweet(c *gin.Context, tweet *model.Tweet) error {
	db, err := connect()
	if err != nil {
		return err
	}

	defer db.Close()

	if err := db.Create(&tweet).Error; err != nil {
		return err
	}

	return nil
}
func (tp *tweetPersistence) GetTweetByID(c *gin.Context, id int64) (*model.Tweet, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	// tweets := []*model.Tweet{}
	tweet := &model.Tweet{}
	// for _, id := range ids {
	// 	if err := db.Where("id = ?", id).Find(&tweet).Error; err != nil {
	// 		return nil, err
	// 	}
	// 	tweets = append(tweets, tweet)
	// 	log.Println(tweet)
	// 	log.Println(&tweets)
	// }

	if err := db.Where("id = ?", id).Find(&tweet).Error; err != nil {
		return nil, err
	}
	return tweet, nil
}
func (tp *tweetPersistence) GetTweetByIDs(c *gin.Context, ids []int64) ([]*model.Tweet, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	tweet := []*model.Tweet{}
	if err := db.Find(&tweet, ids).Error; err != nil {
		return nil, err
	}

	return tweet, nil
}
func (tp *tweetPersistence) DeleteTweetByID(c *gin.Context, id int64) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	tweet := &model.Tweet{}
	if err := db.Model(&tweet).Where("id = ?", id).Update("is_private", true).Error; err != nil {
		return err
	}
	return nil
}
