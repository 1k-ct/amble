package rest

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/1k-ct/amble/app/domain/model"
	"github.com/1k-ct/amble/app/infra/persistence"
	"github.com/1k-ct/amble/app/usecase"
	"github.com/1k-ct/amble/pkg/appErrors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type TweetHandler interface {
	Tweet(c *gin.Context)
	GetTweet(c *gin.Context)
	GetTweets(c *gin.Context)
	UpdateTweet(c *gin.Context)
}

type tweetHandler struct {
	tweetUseCase usecase.TweetUseCase
}

func NewTweetHandler(tu usecase.TweetUseCase) TweetHandler {
	return &tweetHandler{
		tweetUseCase: tu,
	}
}

func (th *tweetHandler) Tweet(c *gin.Context) {

	var request struct {
		UserStaticID string `json:"user_static_id"`
		Content      string `json:"content"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, appErrors.ErrorJSON)
		return
	}
	userName, err := persistence.NewAccountPersistence().GetUserName(request.UserStaticID)
	if err != nil {
		if gorm.ErrRecordNotFound != nil {
			c.JSON(http.StatusBadRequest, appErrors.ErrMeatdataMsg(err, appErrors.ErrRecordDatabase))
			return
		}
		c.JSON(http.StatusInternalServerError, appErrors.ErrMeatdataMsg(err, appErrors.ServerError))
		return
	}
	staticID := uuid.NewV4().String()
	tweet := &model.Tweet{
		ID:             0,
		UserStaticID:   request.UserStaticID,
		StaticID:       staticID,
		IsPrivate:      false,
		Name:           userName,
		Content:        request.Content,
		LikedCount:     0,
		RetweetedCount: 0,
		ReplyCount:     0,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := th.tweetUseCase.RegisterTweet(tweet); err != nil {
		c.JSON(http.StatusInternalServerError, appErrors.ServerError)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"tweet": "ok"})
}
func (th *tweetHandler) GetTweet(c *gin.Context) {
	// "/api/v1/tweet/:id" example uri "/api/v1/tweet/2"
	staticID := c.Param("id")
	tweet, err := th.tweetUseCase.GetTweetByID(staticID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, appErrors.ErrMeatdataMsg(err, appErrors.ErrorJSON))
			return
		}
		c.JSON(http.StatusInternalServerError, appErrors.ServerError)
		return
	}
	c.JSON(http.StatusOK, tweet)
}
func (th *tweetHandler) GetTweets(c *gin.Context) {
	type request struct {
		IDs []int64 `json:"ids"`
	}
	req := &request{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, appErrors.ErrorJSON)
		return
	}
	tweets, err := th.tweetUseCase.GetTweetByIDs(req.IDs)
	if err != nil {
		c.JSON(http.StatusBadRequest, appErrors.ErrorJSON)
		return
	}
	c.JSON(http.StatusOK, tweets)
}
func (th *tweetHandler) UpdateTweet(c *gin.Context) {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, appErrors.ErrorJSON)
		return
	}
	// TODO 更新がすでにされている場合の処理を書く
	// ok, err := delete....(); if err != nil ...{}; if !ok {}
	if err := th.tweetUseCase.DeleteTweetByID(int64(ID)); err != nil {
		c.JSON(http.StatusInternalServerError, appErrors.ServerError)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"msg": "ok"})
}
