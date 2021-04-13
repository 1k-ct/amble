package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/1k-ct/twitter-dem/app/domain/model"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGodotenvLoad(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		log.Println("AccountHandler Router")
		t.Error(err)
	}
	fmt.Println(os.Getenv("SECRET_KEY"))
}
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
	if err := godotenv.Load(); err != nil {
		log.Println("AccountHandler Router")
		log.Fatal(err)
	}
	key := os.Getenv("SECRET_KEY")
	router := InitRouters(key)
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(method, url, body)
	router.ServeHTTP(w, c.Request)
	return
}
func TestExample(t *testing.T) {
	// body := bytes.NewBufferString("")

	w, _ := testform("GET", "/api/v1/tweet/1", nil)

	// var tweetMap map[string]interface{}
	// responseBody := w.Body.String()
	// err := json.Unmarshal([]byte(responseBody), &tweetMap)
	// if err != nil {
	// 	t.Error(err)
	// }
	// fmt.Println(tweetMap["code"])
	// fmt.Println(responseBody)

	assert.Equal(t, 400, w.Code)

	w, _ = testform("GET", "/api/v1/tweet/6", nil)
	assert.Equal(t, 200, w.Code)

}
func TestTweetHandler(t *testing.T) {
	w, _ := testform("GET", "/api/v1/tweet/4", nil)

	assert.Equal(t, 200, w.Code)
	log.Println(w.Body.String())
}
func TestDeleteTweetByID(t *testing.T) {
	w, _ := testform("PUT", "/api/v1/tweet/3", nil)

	assert.Equal(t, 204, w.Code)
}
func TestGetTweets(t *testing.T) {
	var id1 int64 = 6
	var id2 int64 = 7
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
