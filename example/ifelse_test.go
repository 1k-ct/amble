package example

import (
	"errors"
	"fmt"
	"log"
	"testing"
)

func print() string {
	return "a 1 callback function"
}
func TestMain(m *testing.M) {
	songs := []song{
		{Name: "a", No: 1, Func: print},
		{Name: "b", No: 2},
		{Name: "c", No: 3},
		{Name: "d", No: 4},
		{Name: "e", No: 5},
	}
	song, err := findSongNumber(songs, "a")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(song.Name)
	if f, ok := song.Func.(func() string); ok {
		fmt.Println(f())
	}
}

type song struct {
	Name string
	No   int
	Func interface{}
	// Func func()
}

func findSongNumber(songs []song, songName string) (*song, error) {
	for i := 0; i < len(songs); i++ {
		if songs[i].Name == songName {
			return &songs[i], nil
		}
	}
	return nil, errors.New("invalid song")
}
