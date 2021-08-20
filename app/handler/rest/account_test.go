package rest

import (
	"fmt"
	"os"
	"testing"

	"github.com/1k-ct/amble/pkg/appErrors"
	"github.com/joho/godotenv"
)

func TestGoEnv(t *testing.T) {
	if err := godotenv.Load(""); err != nil {
		fmt.Println(appErrors.ErrMeatdataMsg(err, appErrors.ErrInvalidToken))
	}
	fmt.Println(os.Getenv("PORT"))
}
