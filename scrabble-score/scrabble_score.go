package scrabble

import (
	"strings"
)

//Score  compute the scrabble score for the word
func Score(text string) int {
	text = strings.ToLower(text)
	result := 0
	for _, key := range text {
		result += defineValue(key)
	}
	return result
}

func defineValue(letter rune) int {
	switch letter {
	case 'q', 'z':
		return 10
	case 'j', 'x':
		return 8
	case 'k':
		return 5
	case 'f', 'h', 'v', 'w', 'y':
		return 4
	case 'b', 'c', 'm', 'p':
		return 3
	case 'd', 'g':
		return 2
	default:
		return 1
	}
}
