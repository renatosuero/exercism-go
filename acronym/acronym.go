// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	s = strings.ToUpper(s)
	r, _ := regexp.Compile(`\b(\w)`)
	chars := r.FindAllString(s, -1)
	return strings.Join(chars, "")
}
