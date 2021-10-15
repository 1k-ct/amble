package abcrypto

import (
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/xerrors"
)

// type
func GenerateFromBytes(passoword string, cost int) (string, error) {
	if cost < 9 {
		return "", xerrors.New("if cost < 9 is err")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(passoword), cost)
	if err != nil {
		return "", xerrors.Errorf("%v", err)
	}
	return string(hash), nil
}
func CompareHashAndPassword(hashedPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err == bcrypt.ErrHashTooShort {
		return false, xerrors.Errorf("%v", err)
	}
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, xerrors.Errorf("%v", err)
	}

	if err != nil {
		return false, xerrors.Errorf("%v", err)
	}
	return true, nil
}
