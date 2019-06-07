package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

//Valid check if the value is valid based in lugh algorithm
func Valid(text string) bool {
	text = strings.Replace(text, " ", "", -1)
	match, _ := regexp.MatchString(`\D`, text)
	if len(text) < 2 || match {
		return false
	}
	l := len(text) - 1
	for i := l; i > 0; i = i - 2 {
		num, _ := strconv.Atoi(text[i-1 : i])
		num *= 2
		if num > 9 {
			num -= 9
		}
		text = text[:i-1] + strconv.Itoa(num) + text[i:]
	}
	sum := 0
	for _, i := range text {
		add, _ := strconv.Atoi(string(i))
		sum += add
	}
	return sum%10 == 0
}
