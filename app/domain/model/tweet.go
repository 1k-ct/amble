package model

import "time"

type Tweet struct {
	ID           int64  `gorm:"primary_key" json:"id"`
	IsPrivate    bool   `json:"is_private"`
	Name         string `json:"name"`
	Content      string `json:"content"`
	LikeCount    int64  `json:"like_count"`
	RetweetCount int64  `json:"retweet_count"`
	// reply の lenでするかも
	ReplyCount int64     `json:"reply_count"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	// TODO リプライも追加 reply content [] array ? list ?
	// Replies []Reply
	// hang from ぶらさがっているツイート (reply)
}
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type Request struct {
	ID int64
}
