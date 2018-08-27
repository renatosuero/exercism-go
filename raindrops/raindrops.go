package raindrops

import "strconv"

//Convert a number to a string, the contents of which depend on the number's factors.
func Convert(value int) string {
	text := ""
	if value%3 == 0 {
		text += "Pling"
	}
	if value%5 == 0 {
		text += "Plang"
	}
	if value%7 == 0 {
		text += "Plong"
	}
	if text == "" {
		text = strconv.Itoa(value)
	}
	return text
}
