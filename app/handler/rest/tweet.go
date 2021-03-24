package rest

import (
	"net/http"
	"strconv"

	"github.com/1k-ct/twitter-dem/app/domain/model"
	"github.com/1k-ct/twitter-dem/app/usecase"
	"github.com/gin-gonic/gin"
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

type ApplicationError struct {
	Code  int    `json:"code"`
	Level string `json:"level"`
	Msg   string `json:"msg"`
}

var serverError = &ApplicationError{
	Code:  http.StatusInternalServerError,
	Level: "Error",
	Msg:   "An error has occurred inside the server.",
}
var errorJSON = &ApplicationError{
	Code:  http.StatusBadRequest,
	Level: "Error",
	Msg:   "I couldn't read the json.",
}

func (th *tweetHandler) Tweet(c *gin.Context) {

	tweet := &model.Tweet{}
	if err := c.BindJSON(&tweet); err != nil {
		c.JSON(http.StatusBadRequest, errorJSON)
		return
	}

	err := th.tweetUseCase.RegisterTweet(c, tweet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, serverError)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"msg": "ok"})
}
func (th *tweetHandler) GetTweet(c *gin.Context) {
	// "/api/v1/tweet/:id" example uri "/api/v1/tweet/2"
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorJSON)
		return
	}
	tweet, err := th.tweetUseCase.GetTweetByID(c, int64(ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, serverError)
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
		c.JSON(http.StatusBadRequest, errorJSON)
		return
	}
	tweets, err := th.tweetUseCase.GetTweetByIDs(c, req.IDs)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorJSON)
		return
	}
	c.JSON(http.StatusOK, tweets)
}
func (th *tweetHandler) UpdateTweet(c *gin.Context) {
	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorJSON)
		return
	}
	// TODO 更新がすでにされている場合の処理を書く
	// ok, err := delete....(); if err != nil ...{}; if !ok {}
	if err := th.tweetUseCase.DeleteTweetByID(c, int64(ID)); err != nil {
		c.JSON(http.StatusInternalServerError, serverError)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"msg": "ok"})
}
