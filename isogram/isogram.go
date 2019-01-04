package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram check if a word or phrase is an isogram
func IsIsogram(word string) bool {
	if word == "" {
		return true
	}
	word = strings.ToLower(word)
	letters := map[rune]bool{}
	for _, char := range word {
		if !unicode.IsLetter(char) {
			continue
		}
		if letters[char] {
			return false
		}

		letters[char] = true
	}
	return true
}
