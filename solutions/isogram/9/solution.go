package isogram

import (
	"unicode"
)

// IsIsogram checks whether a string is an isogram or not
func IsIsogram(word string) bool {
	letters := map[rune]bool{}

	for _, r := range word {
		if !unicode.IsLetter(r) {
			continue
		}
		r = unicode.ToLower(r)
		if letters[r] {
			return false
		}
		letters[r] = true
	}
	return true
}
