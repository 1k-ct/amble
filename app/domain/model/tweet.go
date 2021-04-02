package model

import "time"

type Tweet struct {
	ID           int64     `gorm:"primary_key" json:"id"`
	StaticID     int64     `json:"static_id"`
	IsPrivate    bool      `json:"is_private"`
	Name         string    `json:"name"`
	Content      string    `json:"content"`
	LikeCount    int64     `json:"like_count"`
	IsLiked      bool      `json:"is_liked"`
	RetweetCount int64     `json:"retweet_count"`
	IsRetweeted  bool      `json:"is_retweeted"` // reply の lenでするかも
	ReplyCount   int64     `json:"reply_count"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	// TODO リプライも追加 reply content [] array ? list ?
	// Replies []Reply
	// hang from ぶらさがっているツイート (reply)
}
type Reply struct {
	ID           int64     `gorm:"primary_key" json:"id"`
	StaticID     int64     `json:"static_id"`
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
	StaticID int64 `json:"static_id"`
}
type LikedUsers struct {
	Users []LikedUser `json:"users"`
}

type RetweetedUser struct {
	StaticID int64 `json:"static_id"`
}
type RetweetedUsers struct {
	Users []RetweetedUser `json:"users"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type Request struct {
	ID int64 `json:"id"`
}
