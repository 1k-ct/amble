package rest

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/1k-ct/twitter-dem/app/domain/model"
	"github.com/1k-ct/twitter-dem/app/usecase"
	"github.com/1k-ct/twitter-dem/pkg/appErrors"
	"github.com/1k-ct/twitter-dem/pkg/jwtToken"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type AccountHandler interface {
	Verify(c *gin.Context)
	SignUp(c *gin.Context)
	Login(c *gin.Context)
	Refresh(c *gin.Context)
}
type accountHandler struct {
	accountUseCase usecase.AccountUseCase
}

func NewAccountHandler(au usecase.AccountUseCase) AccountHandler {
	return &accountHandler{
		accountUseCase: au,
	}
}

func (ah *accountHandler) Verify(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
func (ah *accountHandler) SignUp(c *gin.Context) {
	user := &model.User{}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, appErrors.ErrMeatdataMsg(err, appErrors.ErrorJSON))
		return
	}
	if user.UserName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "UserNameは必須です。"})
		return
	}
	if user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Passwordは必須です。"})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "パスワードに問題があります。"})
	}
	user.Password = string(hash)

	if err := ah.accountUseCase.RegisterUserAccount(user); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, appErrors.ErrRecordDatabase)
			return
		}
		c.JSON(http.StatusInternalServerError, appErrors.ServerError)
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, user)
}
func (ah *accountHandler) Login(c *gin.Context) {
	u := &model.User{}
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, appErrors.ErrMeatdataMsg(err, appErrors.ErrorJSON))
		return
	}
	user, err := ah.accountUseCase.FindByID(u.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, appErrors.ErrMeatdataMsg(err, appErrors.ErrRecordDatabase))
			return
		}
		c.JSON(http.StatusInternalServerError, appErrors.ServerError)
		return
	}
	if err := godotenv.Load(); err != nil {
		log.Println("Login handler")
		log.Fatal(err)
	}
	ts, err := jwtToken.CreateToken(user.ID, user.UserName, os.Getenv("SECRET_KEY"), os.Getenv("REFRESH_KEY"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, appErrors.ErrNotCreateToken)
		return
	}
	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)
}
func (ah *accountHandler) Refresh(c *gin.Context) {
	if err := godotenv.Load(); err != nil {
		log.Println("Refresh handler")
		log.Fatal(err)
	}
	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, appErrors.ErrorJSON)
		return
	}
	refreshToken := mapToken["refresh_token"]

	//verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_KEY")), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Refresh token expired")
		return
	}
	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		userId, roleOk := claims["user_id"].(string)
		if !roleOk {
			c.JSON(http.StatusUnprocessableEntity, "unauthorized")
			return
		}

		user, err := ah.accountUseCase.FindByID(userId)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, "User's not found ")
		}

		ts, createErr := jwtToken.CreateToken(userId, user.UserName, os.Getenv("SECRET_KEY"), os.Getenv("REFRESH_KEY"))
		if createErr != nil {
			c.JSON(http.StatusForbidden, appErrors.ErrMeatdataMsg(createErr, appErrors.ErrNotCreateToken))
			return
		}
		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		c.JSON(http.StatusCreated, tokens)
	} else {
		c.JSON(http.StatusUnauthorized, "refresh expired")
	}
}
