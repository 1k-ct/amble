package router

import (
	"github.com/1k-ct/amble/app/handler/rest"
	"github.com/1k-ct/amble/app/infra/persistence"
	"github.com/1k-ct/amble/app/usecase"
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
		v1.PUT("/tweet/:id", tweetHandler.UpdateTweet)

		v1.POST("/tweets", tweetHandler.GetTweets)
	}

	return r
}
