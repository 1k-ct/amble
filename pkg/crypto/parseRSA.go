package cryptoRSA

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/xerrors"
)

var (
	// "invalid data"
	ErrInvalidData = "invalid data"
	// "invalid key data"
	ErrInvalidKey = "invalid key data"
	// "key cannot decode"
	ErrDecode = "key cannot decode"
	// "type error"
	ErrType = "type error"
)

func PemDecoder(data []byte) (*pem.Block, error) {
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, xerrors.New(ErrInvalidKey)
	}
	return block, nil
}

// create secret key RSA
// openssl genrsa 4096 > secret.key
// create public key RSA
// openssl rsa -pubout < secret.key > public.key
// 秘密鍵の読み込み (RSA)
func ParseRSAPrivateKey(privateKeyData []byte) (*rsa.PrivateKey, error) {
	keyBlock, _ := pem.Decode(privateKeyData)
	if keyBlock == nil {
		return nil, xerrors.New(ErrDecode)
	}
	if keyBlock.Type != "RSA PRIVATE KEY" {
		return nil, xerrors.New(ErrInvalidKey)
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		return nil, xerrors.New(ErrInvalidKey)
	}
	privateKey.Precompute()

	if err := privateKey.Validate(); err != nil {
		return nil, xerrors.New(ErrInvalidKey)
	}
	return privateKey, nil
}

// create secret key RSA
// openssl genrsa 4096 > secret.key
// create public key RSA
// openssl rsa -pubout < secret.key > public.key
// 公開鍵の読み込み (RSA)
func ParseRSAPublicKey(publicKeyData []byte) (*rsa.PublicKey, error) {
	publicKeyBlock, _ := pem.Decode(publicKeyData)
	if publicKeyBlock == nil {
		return nil, xerrors.New(ErrDecode)
	}
	if publicKeyBlock.Type != "PUBLIC KEY" {
		return nil, xerrors.New(ErrInvalidKey)
	}

	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		return nil, errors.New(ErrInvalidKey)
	}
	pub, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, xerrors.Errorf("%v *rsa.PublicKey", ErrType)
	}

	return pub, nil
}

// claims の変更できる
// example generate token and claims
func GenerateToken(userID string, now int64, secretKey interface{}) (string, error) {
	claims := jwt.StandardClaims{
		Subject:   userID,
		IssuedAt:  now,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return tokenString, xerrors.Errorf("failed to sign token : %w", err)
	}

	return tokenString, nil
}

// publicKey = *rsa.PublicKey
// VerifyToken
func VerifyToken(token string, publicKey interface{}) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, xerrors.Errorf("unexpected method. got :%w", method)
		}
		return publicKey, nil
	})
	if err != nil {
		return nil, xerrors.New(err.Error())
	}
	if !parsedToken.Valid {
		return nil, xerrors.New(err.Error())
	}
	return parsedToken, nil
}
