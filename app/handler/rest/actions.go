package rest

import (
	"errors"
	"net/http"
	"time"

	"github.com/1k-ct/twitter-dem/app/domain/model"
	"github.com/1k-ct/twitter-dem/app/infra/persistence"
	"github.com/1k-ct/twitter-dem/app/usecase"
	"github.com/1k-ct/twitter-dem/pkg/appErrors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type ActionsHandler interface {
	Like(c *gin.Context)
	Retweet(c *gin.Context)
	Reply(c *gin.Context)

	GetLikeUser(c *gin.Context)
	GetRetweetUser(c *gin.Context)
	GetReply(c *gin.Context)
}
type actionsHandler struct {
	actionsUseCase usecase.ActionsUseCase
}

func NewActionsHandler(au usecase.ActionsUseCase) ActionsHandler {
	return &actionsHandler{
		actionsUseCase: au,
	}
}
func (ah *actionsHandler) Like(c *gin.Context) {
	var request struct {
		ToTweetID    string `json:"to_tweet_id"`
		UserStaticID string `json:"user_static_id"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, appErrors.ErrMeatdataMsg(err, appErrors.ErrorJSON))
		return
	}
	if err := ah.actionsUseCase.Like(request.UserStaticID, request.ToTweetID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, appErrors.ErrMeatdataMsg(err, appErrors.ErrorJSON))
			return
		}
		c.JSON(http.StatusInternalServerError, appErrors.ServerError)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "like ok"})
}
func (ah *actionsHandler) Retweet(c *gin.Context) {
	var request struct {
		ToTweetID    string `json:"to_tweet_id"`
		UserStaticID string `json:"user_static_id"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, appErrors.ErrMeatdataMsg(err, appErrors.ErrorJSON))
		return
	}
	if err := ah.actionsUseCase.Retweet(request.UserStaticID, request.ToTweetID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, appErrors.ErrMeatdataMsg(err, appErrors.ErrorJSON))
			return
		}
		c.JSON(http.StatusInternalServerError, appErrors.ServerError)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "retwee ok"})
}
func (ah *actionsHandler) Reply(c *gin.Context) {
	var request struct {
		ToTweetID    string `json:"to_tweet_id"`
		UserStaticID string `json:"user_static_id"`
		Content      string `json:"content"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, appErrors.ErrMeatdataMsg(err, appErrors.ErrorJSON))
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
	reply := &model.Reply{
		ToTweetID:    request.ToTweetID,
		UserStaticID: request.UserStaticID,
		StaticID:     staticID,
		IsPrivate:    false,
		Name:         userName,
		Content:      request.Content,
		LikeCount:    0,
		IsLiked:      false,
		RetweetCount: 0,
		IsRetweeted:  false,
		ReplyCount:   0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err := ah.actionsUseCase.Reply(reply); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, appErrors.ErrMeatdataMsg(err, appErrors.ErrorJSON))
			return
		}
		c.JSON(http.StatusInternalServerError, appErrors.ServerError)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "retwee ok"})
}
func (ah *actionsHandler) GetLikeUser(c *gin.Context) {
	toTweetID := c.Param(":id")
	likedUsers, err := ah.actionsUseCase.GetLikeUser(toTweetID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, appErrors.ErrMeatdataMsg(err, appErrors.ErrorJSON))
			return
		}
		c.JSON(http.StatusInternalServerError, appErrors.ServerError)
		return
	}
	// type likedUser struct {
	// 	ToTweetID    string `json:"to_tweet_id"`
	// 	UserStaticID string `json:"user_static_id"`
	// }
	// type response struct {
	// 	LikedUsers []*likedUser
	// }
	// likes := []*likedUser{}
	// for _, lu := range likedUsers {
	// 	liked := &likedUser{}

	// 	liked.ToTweetID = lu.ToTweetID
	// 	liked.UserStaticID = lu.UserStaticID
	// 	likes = append(likes, liked)
	// }
	// res := &response{
	// 	LikedUsers: likes,
	// }

	c.JSON(http.StatusOK, gin.H{"response": likedUsers})
}
func (ah *actionsHandler) GetRetweetUser(c *gin.Context) {
	toTweetID := c.Param(":id")
	retweetUsers, err := ah.actionsUseCase.GetRetweetUser(toTweetID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, appErrors.ErrMeatdataMsg(err, appErrors.ErrorJSON))
			return
		}
		c.JSON(http.StatusInternalServerError, appErrors.ServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": retweetUsers})
}
func (ah *actionsHandler) GetReply(c *gin.Context) {
	toTweetID := c.Param(":id")
	replies, err := ah.actionsUseCase.GetReply(toTweetID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, appErrors.ErrMeatdataMsg(err, appErrors.ErrorJSON))
			return
		}
		c.JSON(http.StatusInternalServerError, appErrors.ServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": replies})
}
