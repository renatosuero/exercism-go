package scrabble

import (
	"strings"
)


func Score(text string) int {
	text = strings.ToLower(text)
	result :=0
	for _, key := range text {
		letter := string(key)
		result+= defineValue(letter)
	}
	return result
}

func defineValue(letter string) int{
		if letter == "q" || letter == "z" {
			return 10
		}
		if letter == "j" || letter == "x" {
			return 8
		}
		if letter == "k" {
			return 5
		}
		if letter == "f" || letter == "h" || letter == "v" || letter == "w" || letter == "y" {
			return 4
		}
		if letter == "b" || letter == "c" || letter == "m" || letter == "p" {
			return 3
		}
		if letter == "d" || letter == "g" {
		return 2	
		}
		return 1
}
