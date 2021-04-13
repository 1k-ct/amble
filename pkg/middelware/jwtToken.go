// Copyright (c) 2020 Steven Victor
// https://github.com/tienbm90/simple-jwt-auth/blob/master/LICENSE
package middelware

import (
	"net/http"

	"github.com/1k-ct/twitter-dem/pkg/appErrors"
	"github.com/1k-ct/twitter-dem/pkg/jwtToken"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwtToken.TokenValid(c.Request, secretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, appErrors.ErrMeatdataMsg(err, appErrors.ErrInvalidToken))
			c.Abort()
			return
		}
		c.Next()
	}
}
