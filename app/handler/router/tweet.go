package router

import (
	"github.com/1k-ct/twitter-dem/app/handler/rest"
	"github.com/1k-ct/twitter-dem/app/infra/persistence"
	"github.com/1k-ct/twitter-dem/app/usecase"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.Default()
	apiGroup := r.Group("/api")
	{
		TweetHandler(apiGroup)
	}

	return r
}
func TweetHandler(r *gin.RouterGroup) (R gin.IRoutes) {
	tweetPersistence := persistence.NewTweetPersistence()
	tweetUseCase := usecase.NewTweetUseCase(tweetPersistence)
	tweetHandler := rest.NewTweetHandler(tweetUseCase)

	v1 := r.Group("/v1")
	{
		v1.POST("/tweet", tweetHandler.Tweet)
		v1.GET("/tweet/:id", tweetHandler.GetTweet)
	}

	return r
}
