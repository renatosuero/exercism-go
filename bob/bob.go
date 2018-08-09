// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)

const (
	fine     = "Fine. Be that way!"
	whoa     = "Whoa, chill out!"
	sure     = "Sure."
	whatever = "Whatever."
	calm     = "Calm down, I know what I'm doing!"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	isQuestion := checkQuestion(remark)
	isUpper := checkToUpper(remark)
	if isQuestion && isUpper {
		return calm
	}
	if isQuestion {
		return sure
	}
	if isUpper {
		return whoa
	}
	if remark == "" {
		return fine
	}
	return whatever
}
func checkQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}
func checkToUpper(remark string) bool {
	return strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark
}
