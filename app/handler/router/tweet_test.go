package router

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func testRouter(
	InitRouter func() *gin.Engine, cors, uri string, body io.Reader,
) (*httptest.ResponseRecorder, error) {
	router := InitRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest(cors, uri, body)
	if err != nil {
		return nil, err
	}
	router.ServeHTTP(w, req)
	return w, nil
}
func TestExample(t *testing.T) {
	// body := bytes.NewBufferString("")
	w, err := testRouter(InitRouters, "GET", "/api/v1/tweet/3", nil /*ここにbody*/)
	if err != nil {
		return
	}
	assert.Equal(t, 200, w.Code)
}
func TestTweetHandler(t *testing.T) {
	router := InitRouters()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/v1/tweet/4", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	log.Println(w.Body.String())
}
