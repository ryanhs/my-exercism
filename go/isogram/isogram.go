package isogram

import "strings"

// ContainsCount tells whether a contains how many x.
func ContainsCount(a []string, x string) int {
	c := 0

	for _, n := range a {
		if x == n {
			c++
		}
	}

	return c
}

func IsIsogram(word string) bool {
	str := strings.ToLower(word)
	split := strings.Split(str, "")

	for _, ch := range split {
		if ch == " " || ch == "-" {
			continue
		}

		count := ContainsCount(split, ch)
		if count > 1 {
			return false
		}
	}

	return true
}
