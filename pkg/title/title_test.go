package title_test

import (
	"fmt"
	"testing"

	"github.com/1k-ct/twitter-dem/pkg/title"
)

func TestFetchVideoInfo(t *testing.T) {
	url := "Ct6BUPvE2sM"
	data, err := title.FetchVideoInfo(url)
	if err != nil {
		t.Error(err)
	}
	res := title.ResponsJson{
		Title:        data.Title,
		AuthorName:   data.AuthorName,
		AuthorURL:    data.AuthorURL,
		Type:         data.Type,
		Version:      data.Version,
		ProviderName: data.ProviderName,
		ProviderURL:  data.ProviderURL,
		ThumbnailURL: data.ThumbnailURL,
	}
	fmt.Println(res)
	// res := `{PIKOTARO - PPAP (Pen Pineapple Apple Pen) (Long Version) [Official Video] Ultra Music https://www.youtube.com/c/Ultramusic video 1.0 YouTube https://www.youtube.com/ https://i.ytimg.com/vi/Ct6BUPvE2sM/hqdefault.jpg}`
}
