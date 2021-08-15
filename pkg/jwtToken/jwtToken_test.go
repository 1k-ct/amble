package jwtToken

import (
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
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

func TestUUID(t *testing.T) {
	u0 := uuid.NewV4()
	fmt.Println("u1:", u0)
	u1 := uuid.NewV4().String()
	fmt.Println("u1:", u1)

	u2, err := uuid.FromString("{1831de38-08ee-448e-8f14-9d70357ab3ec}")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("u2:", u2)
}

func TestCreateToken(t *testing.T) {
	token := testCreateToken()
	if token == "" {
		t.Error("fatal create token")
	}
	fmt.Println(token)
}
func TestTokenValid(t *testing.T) {
	token := testCreateToken()
	if token == "" {
		t.Error("fatal create token")
	}

	accessToken := fmt.Sprintf(`"access_token": %v`, token)
	h := &http.Request{
		Header: map[string][]string{"Authorization": {accessToken}},
	}

	jwtToken, err := VerifyToken(h, secretKey)
	if err != nil {
		t.Error("error verify token")
	}
	if !jwtToken.Valid {
		t.Error("error token valid")
	}

	if err := TokenValid(h, secretKey); err != nil {
		t.Error("error verify token")
	}
	fmt.Println("token ok!")
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
	fmt.Println(token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.NotZero(t, payload.ID)
	require.Equal(t, userName, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}
