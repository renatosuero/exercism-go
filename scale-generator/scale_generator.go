package scale

import (
	"strings"
)

// Scale generate the musical scale starting with the tonic and following the specified interval pattern.
func Scale(tonic, interval string) []string {
	scale := checkTonic(tonic)
	final := []string{}
	tonic = strings.Title(tonic)
	for index, note := range scale {
		if tonic == note {
			final = append(final, scale[index:]...)
			final = append(final, scale[0:index]...)
			continue
		}
	}
	x := []string{}
	if interval == "" {
		return final
	}
	i := 0
	for _, note := range []rune(interval) {
		x = append(x, final[i])
		switch note {
		case 'm':
			i++
		case 'M':
			i += 2
		case 'A':
			i += 3
		}
	}
	return x
}
func checkTonic(tonic string) []string {
	switch tonic {
	case "C", "a", "G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#":
		return []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		return []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
	}
	return []string{}
}
