package abcrypto

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHash(t *testing.T) {
	pass := []byte("hello world\n")
	var sum [32]byte
	sum = sha256.Sum256(pass)
	for i := 0; i < 32768; i++ {
		sum = sha256.Sum256(sum[:])
		// fmt.Printf("%x\n", sum)
	}
	fmt.Printf("result: %x\n", sum)
}

func TestGenerateFromPassword(t *testing.T) {
	password := []byte("hello world")
	hashed, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(string(hashed))
}

func TestComppareHashAndPasswords(t *testing.T) {
	password := []byte("hello world")
	// hashed, err := bcrypt.GenerateFromPassword(password, 15)
	// if err != nil {
	// 	t.Fatal(err.Error())
	// }
	hashed := "$2a$10$rlCC5XO0wkZejJqEZa.cd.MxaARzp9bYM2PGs8d/PPz04N4xCkcZ."
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), password); err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println("ok")

	hashed1 := "$2a$10$rlCC5XO0wkZejJqEZa.cd.MxaARzp9bYM2PGs8d/PPz04N4xCkc"
	err := bcrypt.CompareHashAndPassword([]byte(hashed1), password)
	if err != nil {
		fmt.Println("hashed error ok")
	}
	if err != bcrypt.ErrHashTooShort {
		t.Fatal(err.Error())
	}
	fmt.Println("ok")

	falsePassword := []byte("falsePassword")
	hashed2 := "$2a$10$rlCC5XO0wkZejJqEZa.cd.MxaARzp9bYM2PGs8d/PPz04N4xCkcZ."
	err = bcrypt.CompareHashAndPassword([]byte(hashed2), falsePassword)
	if err != nil {
		// t.Fatal(err.Error())
		fmt.Println("password false error ok")
	}
	if err != bcrypt.ErrMismatchedHashAndPassword {
		t.Fatal(err.Error())
	}
	cost, err := bcrypt.Cost([]byte(hashed))
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(cost)
}
