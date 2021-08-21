package title_test

import (
	"testing"

	"github.com/1k-ct/amble/pkg/title"
	"github.com/stretchr/testify/require"
)

func TestFetchVideoInfo(t *testing.T) {
	url := "Ct6BUPvE2sM"
	data, err := title.FetchVideoInfo(url)

	require.Nil(t, err, "fetch video info err is not nil")
	require.NotEmpty(t, data, "data")

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
	require.NotEmpty(t, res, "res is not empty")
	// res := `{PIKOTARO - PPAP (Pen Pineapple Apple Pen) (Long Version) [Official Video] Ultra Music https://www.youtube.com/c/Ultramusic video 1.0 YouTube https://www.youtube.com/ https://i.ytimg.com/vi/Ct6BUPvE2sM/hqdefault.jpg}`
}
