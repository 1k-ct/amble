package router

import (
	"github.com/1k-ct/twitter-dem/app/handler/rest"
	"github.com/1k-ct/twitter-dem/app/infra/persistence"
	"github.com/1k-ct/twitter-dem/app/usecase"
	"github.com/1k-ct/twitter-dem/pkg/middelware"
	"github.com/gin-gonic/gin"
)

func AccountHandler(r *gin.RouterGroup, secretKey string) (R gin.IRoutes) {
	accountPersistence := persistence.NewAccountPersistence()
	accountUseCase := usecase.NewAccountUseCase(accountPersistence)
	accountHandler := rest.NewAccountHandler(accountUseCase)

	// if err := godotenv.Load(); err != nil {
	// 	log.Println("AccountHandler Router")
	// 	log.Fatal(err)
	// }
	// secretKey := os.Getenv("SECRET_KEY")
	account := r.Group("/account")
	{
		account.POST("/sign", accountHandler.SignUp)
		account.POST("/login", accountHandler.Login)

		account.GET("/verify", middelware.TokenAuthMiddleware(secretKey), accountHandler.Verify)
		account.POST("/refresh", accountHandler.Refresh)
		account.POST("/profile", middelware.TokenAuthMiddleware(secretKey), accountHandler.EditUserProfile)
	}
	return r
}
