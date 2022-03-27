package list

import "golang.org/x/exp/slices"

func isDupes[E comparable](s []E) bool {
	return len(DupesSlice(s)) != 0
}

func DupesSlice[E comparable](s []E) []E {
	dupes := []E{}
	for i, v := range s {
		if slices.Contains(s[i+1:], v) {
			dupes = append(dupes, v)
		}
	}
	return dupes
}
