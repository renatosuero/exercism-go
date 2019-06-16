package luhn

import (
	"strconv"
	"strings"
)

//Valid check if the value is valid based in lugh algorithm
func Valid(text string) bool {
	text = strings.ReplaceAll(text, " ", "")
	_, err := strconv.Atoi(text)
	if len(text) < 2 || err != nil {
		return false
	}

	sum := 0
	double := len(text)%2 == 0
	for _, r := range text {
		digit := int(r - '0')
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}

		}
		double = !double
		sum += digit
	}
	return sum%10 == 0
}
