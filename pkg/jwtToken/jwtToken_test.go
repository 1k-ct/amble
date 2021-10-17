package jwtToken

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	privKeyPath = "keys/app.rsa"
	pubKeyPath  = "keys/app.rsa.pub"
	userID      = "testID"
	userName    = "testName"
	secretKey   = "testSecretKey"
	refreshKey  = "testRefreshKey"
)

// func TestUUID(t *testing.T) {
// 	u0 := uuid.NewV4()
// 	require.NotEmpty(t, u0, "u0:")
// 	u1 := uuid.NewV4().String()
// 	require.NotEmpty(t, u1, "u1:")

// 	input := "{1831de38-08ee-448e-8f14-9d70357ab3ec}"
// 	u2, err := uuid.FromString(input)
// 	require.Nil(t, err, "u2:")
// 	require.NotEmpty(t, u2, "u2:")

// 	u3, err := uuid.FromString(input + "1")
// 	require.NotNil(t, err, "u3:")
// 	require.NotEmpty(t, u3, "u3:")
// }

func TestCreateToken(t *testing.T) {
	token := testCreateToken()
	require.NotEmpty(t, token, "fatal create token:")
}
func TestTokenValid(t *testing.T) {
	token := testCreateToken()
	require.NotEmpty(t, token, "fatal create token:")

	accessToken := fmt.Sprintf(`"access_token": %v`, token)
	h := &http.Request{
		Header: map[string][]string{"Authorization": {accessToken}},
	}

	jwtToken, err := VerifyToken(h, secretKey)
	require.Nil(t, err, "jwt token not nil")
	require.Equal(t, jwtToken.Valid, true, "valid not")

	err = TokenValid(h, secretKey)
	require.NoError(t, err, "err")
}
func testCreateToken() string {
	got, err := CreateToken(userID, userName, secretKey, refreshKey)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	return got.AccessToken
}

func TestJWTMaker(t *testing.T) {
	secretKey := "1234567890123456789012345678901234567890"
	maker, err := NewJWTMaker(secretKey)
	require.NoError(t, err)

	userName := "UserName"
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(userName, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.NotZero(t, payload.ID)
	require.Equal(t, userName, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}
func TestExtract(t *testing.T) {
	claims := jwt.MapClaims{
		"access_uuid": "1234",
		"user_id":     "12345",
		"user_name":   "123456",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Valid = true

	accessDetails, err := Extract(token)
	if err != nil {
		t.Errorf("%+v\n", err)
	}

	require.NotEmpty(t, accessDetails, "accessDetails:")

	expectedActualAccessDetails := &AccessDetails{
		TokenUuid: "1234",
		UserId:    "12345",
		UserName:  "123456",
	}
	require.Equal(t, accessDetails, expectedActualAccessDetails)

	// error unauthorized
	unauthorizedClaims := jwt.MapClaims{
		"access_uuid": nil,
	}
	unauthorizedToken := jwt.NewWithClaims(jwt.SigningMethodES256, unauthorizedClaims)
	unauthorizedToken.Valid = true

	unauthorizedAccessDetails, err := Extract(unauthorizedToken)
	if assert.Error(t, err, "Extract error:") {
		assert.Equal(t, errors.New("unauthorized"), err)
	}
	require.Nil(t, unauthorizedAccessDetails, "accessDetails:")

	// something went wrong
	wrongToken := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"access_uuid": "1234",
		"user_id":     "12345",
		"user_name":   "123456",
	})
	wrongToken.Valid = false
	wrongAccessDetails, err := Extract(wrongToken)
	if assert.Error(t, err, "Extract error:") {
		assert.Equal(t, errors.New("something went wrong"), err)
	}
	require.Nil(t, wrongAccessDetails, "wrongAccessDetails:")
}
