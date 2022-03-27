package list

import (
	"reflect"
	"testing"
)

var indexTests = []struct {
	name  string
	s     []int
	want  []int
	want2 bool
}{
	{
		name:  "non dupes slice",
		s:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		want:  []int{},
		want2: false,
	}, {
		name:  "dupes slice",
		s:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 12, 22, 0},
		want:  []int{1, 2, 3, 0},
		want2: true,
	},
}

func TestDupesSlice(t *testing.T) {
	for _, tt := range indexTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DupesSlice(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DupesSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isDupes(t *testing.T) {
	for _, tt := range indexTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDupes(tt.s); got != tt.want2 {
				t.Errorf("isDupes() = %v, want %v", got, tt.want)
			}
		})
	}
}
