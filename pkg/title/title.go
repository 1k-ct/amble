package title

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type VideoInfo struct {
	Title           string `json:"title"`
	AuthorName      string `json:"author_name"`
	AuthorURL       string `json:"author_url"`
	Type            string `json:"type"`
	Height          int    `json:"height"`
	Width           int    `json:"width"`
	Version         string `json:"version"`
	ProviderName    string `json:"provider_name"`
	ProviderURL     string `json:"provider_url"`
	ThumbnailHeight int    `json:"thumbnail_height"`
	ThumbnailWidth  int    `json:"thumbnail_width"`
	ThumbnailURL    string `json:"thumbnail_url"`
	HTML            string `json:"html"`
}
type ResponsJson struct {
	Title        string `json:"title"`
	AuthorName   string `json:"author_name"`
	AuthorURL    string `json:"author_url"`
	Type         string `json:"type"`
	Version      string `json:"version"`
	ProviderName string `json:"provider_name"`
	ProviderURL  string `json:"provider_url"`
	ThumbnailURL string `json:"thumbnail_url"`
}

func FetchVideoInfo(VideoURL string) (*VideoInfo, error) {
	url := "https://www.youtube.com/oembed?format=json&url=https://www.youtube.com/watch?v=" + VideoURL
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	jsonBytes := ([]byte)(byteArray)
	data := new(VideoInfo)

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		return nil, err
	}
	return data, nil
}
