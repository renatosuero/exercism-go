package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

//Valid check if the value is valid based in lugh algorithm
func Valid(text string) bool {
	text = strings.Replace(text, " ", "", -1)
	if len(text) < 2 || checkSpecialChar(text) {
		return false
	}
	text = luhnFormula(text)
	return sumAllValues(text)%10 == 0
}

//checkSpecialChar check if the text contains special chars
func checkSpecialChar(text string) bool {
	match, _ := regexp.MatchString(`\D`, text)
	return match
}

//luhnFormula apply luhn formula
func luhnFormula(value string) string {
	l := len(value) - 1
	for i := l; i > 0; i = i - 2 {
		num, _ := strconv.Atoi(value[i-1 : i])
		num *= 2
		num = isGreaterNine(num)
		value = value[:i-1] + strconv.Itoa(num) + value[i:]
	}
	return value
}

//isGreaterNine if number is greater than 9 subtract 9
func isGreaterNine(number int) int {
	if number > 9 {
		number -= 9
	}
	return number
}

//sumAllValues converts strint into number and sum the value
func sumAllValues(values string) int {
	sum := 0
	for _, i := range values {
		add, _ := strconv.Atoi(string(i))
		sum += add
	}
	return sum
}
