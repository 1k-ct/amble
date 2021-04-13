package jwtToken

import (
	"fmt"
	"testing"

	uuid "github.com/satori/go.uuid"
)

const (
	privKeyPath = "keys/app.rsa"
	pubKeyPath  = "keys/app.rsa.pub"
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
