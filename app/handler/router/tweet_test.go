package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/1k-ct/twitter-dem/app/domain/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func testRouter(
	InitRouter func() *gin.Engine, crud, uri string, body io.Reader,
) (*httptest.ResponseRecorder, error) {
	router := InitRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest(crud, uri, body)
	if err != nil {
		return nil, err
	}
	router.ServeHTTP(w, req)
	return w, nil
}
func testform(method string, url string, body io.Reader) (w *httptest.ResponseRecorder, c *gin.Context) {
	router := InitRouters()
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(method, url, body)
	router.ServeHTTP(w, c.Request)
	return
}
func TestExample(t *testing.T) {
	// body := bytes.NewBufferString("")
	w, err := testRouter(InitRouters, "GET", "/api/v1/tweet/3", nil /*ここにbody*/)
	if err != nil {
		t.Error(err)
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
func TestDeleteTweetByID(t *testing.T) {
	w, _ := testform("PUT", "/api/v1/tweet/3", nil)

	assert.Equal(t, 204, w.Code)
}
func TestGetTweets(t *testing.T) {
	var id1 int64 = 1
	var id2 int64 = 2
	postBody := fmt.Sprintf(`{"ids": [%d, %d]}`, id1, id2)
	body := bytes.NewBufferString(postBody)
	w, _ := testform("POST", "/api/v1/tweets", body)

	assert.Equal(t, 200, w.Code)

	var tweets []model.Tweet

	tweet := w.Body.String()
	byteTweet := []byte(tweet)
	json.Unmarshal(byteTweet, &tweets)

	// fmt.Println(tweets) or fmt.Println(w.Body.String())

	if tweets[0].ID != id1 {
		t.Error("code error")
	}
	if tweets[1].ID != id2 {
		t.Error("code error")
	}
}
