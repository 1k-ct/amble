package cryptoRSA_test

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"testing"
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
