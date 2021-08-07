// Copyright (c) 2020 Steven Victor
// https://github.com/tienbm90/simple-jwt-auth/blob/master/LICENSE
package jwtToken

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

type AccessDetails struct {
	TokenUuid string
	UserId    string
	UserName  string
}
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	TokenUuid    string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

func NewTokenDetails() *TokenDetails {
	return &TokenDetails{}
}
func CreateToken(userID, userName, secretKey, refreshKey string) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 30).Unix() //expires after 30 min
	td.TokenUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = td.TokenUuid + "++" + userID
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["access_uuid"] = td.TokenUuid
	atClaims["user_id"] = userID
	atClaims["user_name"] = userName
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	// td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	td.AccessToken, err = at.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = td.TokenUuid + "++" + userID

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userID
	rtClaims["user_name"] = userName
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	// td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	td.RefreshToken, err = rt.SignedString([]byte(refreshKey))
	if err != nil {
		return nil, err
	}
	return td, nil
}
func TokenValid(r *http.Request, secretKey string) error {
	token, err := VerifyToken(r, secretKey)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}
func VerifyToken(r *http.Request, secretKey string) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// return []byte(os.Getenv("ACCESS_SECRET")), nil
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

//get the token from the request body
func ExtractToken(r *http.Request) string {
	// Authorization "access_token": ey....xxx
	// tokenの　"　を含んでいるか気をつける
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// Extract 抽出
func Extract(token *jwt.Token) (*AccessDetails, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		userId, userOk := claims["user_id"].(string)
		userName, userNameOk := claims["user_name"].(string)
		if !ok || !userOk || !userNameOk {
			return nil, errors.New("unauthorized")
		}
		return &AccessDetails{
			TokenUuid: accessUuid,
			UserId:    userId,
			UserName:  userName,
		}, nil

	}
	return nil, errors.New("something went wrong")
}

// ExtractTokenMetadata トークンのメタデータを抽出
func ExtractTokenMetadata(r *http.Request, secretKey string) (*AccessDetails, error) {
	token, err := VerifyToken(r, secretKey)
	if err != nil {
		return nil, err
	}
	acc, err := Extract(token)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
