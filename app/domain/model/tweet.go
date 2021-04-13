package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Tweet struct {
	ID             int64     `json:"id"`
	UserStaticID   string    `json:"user_static_id"`
	StaticID       string    `json:"static_id"`
	IsPrivate      bool      `json:"is_private"`
	Name           string    `json:"name"`
	Content        string    `json:"content"`
	LikedCount     int64     `json:"liked_count"`
	RetweetedCount int64     `json:"retweeted_count"`
	ReplyCount     int64     `json:"reply_count"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	// TODO リプライも追加 reply content [] array ? list ?
	// Replies []Reply
	// hang from ぶらさがっているツイート (reply)
}
type Reply struct {
	ToTweetID    string    `json:"to_tweet_id"`
	UserStaticID string    `json:"user_static_id"`
	StaticID     string    `json:"static_id"`
	IsPrivate    bool      `json:"is_private"`
	Name         string    `json:"name"`
	Content      string    `json:"content"`
	LikeCount    int64     `json:"like_count"`
	IsLiked      bool      `json:"is_liked"`
	RetweetCount int64     `json:"retweet_count"`
	IsRetweeted  bool      `json:"is_retweeted"`
	ReplyCount   int64     `json:"reply_count"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type LikedUser struct {
	gorm.Model
	ToTweetID    string `json:"to_tweet_id"`
	UserStaticID string `json:"user_static_id"`
}

type RetweetedUser struct {
	gorm.Model
	ToTweetID    string `json:"to_tweet_id"`
	UserStaticID string `json:"user_static_id"`
}

type Response struct {
	Code    int
	Message string
}
type Request struct {
	ID int64
}
