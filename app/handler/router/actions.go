package router

import (
	"github.com/1k-ct/twitter-dem/app/handler/rest"
	"github.com/1k-ct/twitter-dem/app/infra/persistence"
	"github.com/1k-ct/twitter-dem/app/usecase"
	"github.com/gin-gonic/gin"
)

func ActionsHandler(r *gin.RouterGroup, secretKey string) (R gin.IRoutes) {
	actionsPersistence := persistence.NewActionsPersistence()
	actionsUseCase := usecase.NewActionsUseCase(actionsPersistence)
	actionsHandler := rest.NewActionsHandler(actionsUseCase)

	actions := r.Group("/actions")
	// TODO
	// actions.Use(middelware.TokenAuthMiddleware(secretKey))
	{
		actions.POST("/like", actionsHandler.Like)
		actions.POST("/retweet", actionsHandler.Retweet)
		actions.POST("/reply", actionsHandler.Reply)

		// TODO
		// 認証追加
		// 上のかこの下のどちらか
		// account.GET("/verify", middelware.TokenAuthMiddleware(secretKey), accountHandler.Verify)

		// この下の処理は、tweet に含むかも

		actions.GET("/like/:id", actionsHandler.GetLikeUser)
		// actions.PUT("/like",actionsHandler.) delete like

		actions.GET("/retweet/:id", actionsHandler.GetRetweetUser)
		// actions.PUT("/retweet",actionsHandler.) delete retweet

		actions.GET("/reply/:id", actionsHandler.GetReply)
		// actions.PUT("/reply",actionsHandler.) delete reply
	}
	return r
}
