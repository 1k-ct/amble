package main

import (
	"fmt"
	"log"

	"github.com/1k-ct/amble/pkg/abcrypto"
)

func main() {
	password := "abcdefghijklmnopqrstuvwxyz"
	h, err := abcrypto.GenerateFromBytes(password, 10)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	ok, err := abcrypto.CompareHashAndPassword(h, password)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	if !ok {
		log.Fatal("password error")
	}
	fmt.Println("ok")
}
