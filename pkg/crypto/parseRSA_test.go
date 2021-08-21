package cryptoRSA_test

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	cryptoRSA "github.com/1k-ct/amble/pkg/crypto"
	"github.com/stretchr/testify/require"
)

// example
func TestGenerateKey(t *testing.T) {
	// 秘密鍵の生成
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	// fmt.Println(privateKey)
	// 公開鍵の生成
	publicKey := privateKey.PublicKey
	// 暗号化
	secretMessage := []byte("send reinforcements, we're going to advance")
	label := []byte("orders")
	// fmt.Printf("Ciphertext: %x\n", ciphertext)
	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&publicKey,
		secretMessage,
		label)
	if err != nil {
		t.Error(err)
	}
	// 復号
	decryptedBytes, err := privateKey.Decrypt(rand.Reader, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256, Label: label})
	if err != nil {
		t.Error(err)
	}
	// 一応確認
	if !bytes.Equal(secretMessage, decryptedBytes) {
		t.Error(err)
	}
	fmt.Println("decrypted message: ", string(decryptedBytes))
}

func TestParseRSAPrivateKey(t *testing.T) {
	filename := "../../test/private.pem"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("%+v\n", err)
	}
	privateKey, err := cryptoRSA.ParseRSAPrivateKey(data)
	if err != nil {
		t.Errorf("%+v\n", err)
	}

	if reflect.DeepEqual(privateKey, data) {
		t.Errorf("ParseRSAPrivateKey() = %v, want %v", privateKey, data)
	}
}

// https://www.systutorials.com/how-to-generate-rsa-private-and-public-key-pair-in-go-lang/
func TestGenerateRSAKey(t *testing.T) {
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Errorf("%+v\n", err)
	}
	require.NoErrorf(t, err, "%+v\n", err)
	publickey := &privatekey.PublicKey

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privatekey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	privatePem, err := os.Create("./../../test/private.pem")
	require.NoErrorf(t, err, "%+v\n", err)

	err = pem.Encode(privatePem, privateKeyBlock)
	require.NoErrorf(t, err, "%+v\n", err)

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publickey)
	require.NoErrorf(t, err, "%+v\n", err)

	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	publicPem, err := os.Create("./../../test/public.pem")
	require.NoErrorf(t, err, "%+v\n", err)

	err = pem.Encode(publicPem, publicKeyBlock)
	require.NoErrorf(t, err, "%+v\n", err)
}
