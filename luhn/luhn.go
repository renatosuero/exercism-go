package luhn

import (
	"strconv"
	"strings"
)

//Valid check if the value is valid based in lugh algorithm
func Valid(text string) bool {
	text = strings.Replace(text, " ", "", -1)
	_, err := strconv.Atoi(text)
	if len(text) < 2 || err != nil {
		return false
	}

	sum := 0
	for i := 0; i < len(text); i++ {
		digit, _ := strconv.Atoi(text[i : i+1])

		if (len(text)-i)%2 == 0 {
			double := digit * 2
			sum += double/10 + double%10
			continue
		}
		sum += digit
	}
	return sum%10 == 0
}
