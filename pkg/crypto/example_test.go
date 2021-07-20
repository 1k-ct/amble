package cryptoRSA_test

import (
	"fmt"
	"io/ioutil"
	"log"

	cryptoRSA "github.com/1k-ct/twitter-dem/pkg/crypto"
)

// https://pkg.go.dev/crypto/rsa#PublicKey

// A PublicKey represents the public part of an RSA key.
// type PublicKey struct {
// 	N *big.Int // modulus
// 	E int      // public exponent
// }

func exampleParseKeyData() {
	filename := "./publick.key"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	publicKey, err := cryptoRSA.ParseRSAPublicKey(data)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	// *rsa.PublicKey
	fmt.Println(publicKey)
}

// https://pkg.go.dev/crypto/rsa#PrivateKey

// A PrivateKey represents an RSA key
// type PrivateKey struct {
// 	PublicKey            // public part.
// 	D         *big.Int   // private exponent
// 	Primes    []*big.Int // prime factors of N, has >= 2 elements.

// 	// Precomputed contains precomputed values that speed up private
// 	// operations, if available.
// 	Precomputed PrecomputedValues
// }

func exampleParsePrivateKey() {
	filename := "./private.key"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	privateKey, err := cryptoRSA.ParseRSAPrivateKey(data)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	// *rsa.PrivateKey
	fmt.Println(privateKey)
}

func exampleMain() {
	// 	PublicKey の読み込み
	filename := "./private.key"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	// private keyのパース
	privateKey, err := cryptoRSA.ParseRSAPrivateKey(data)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	// jwt の作成
	token, err := cryptoRSA.GenerateToken("123", 12345, privateKey)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	// 	PrivateKey の読み込み
	publicKeyData, err := ioutil.ReadFile("./public.key")
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	// public keyのパース
	pub, err := cryptoRSA.ParseRSAPublicKey(publicKeyData)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	// 生成された署名の検証
	t, err := cryptoRSA.VerifyToken(token, pub)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	fmt.Println(t.Valid)
}
