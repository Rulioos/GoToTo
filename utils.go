package main

import (
	"strings"
	"unicode"
)

//Removing spaces from string
func SpaceStringsBuilder(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

//Transform a list to a set
func SetifyString(s *[]string) {
	m := make(map[string]bool)
	for _, r := range *s {
		m[r] = true
	}

	*s = make([]string, 0)
	for k := range m {
		*s = append(*s, k)
	}

}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
