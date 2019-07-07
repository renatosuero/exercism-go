// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb generates a proverb based in the word list
func Proverb(rhyme []string) []string {
	total := len(rhyme)
	if total < 1 {
		return []string{}
	}
	if total < 2 {
		return []string{"And all for the want of a nail."}
	}
	phrase := make([]string, total)
	for i := 0; i < total-1; i++ {
		phrase[i] = "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."
	}
	phrase[total-1] = "And all for the want of a " + rhyme[0] + "."
	return phrase
}

// Proverb generates a proverb based in the word list
func Proverb2(rhyme []string) []string {
	total := len(rhyme)
	if total < 1 {
		return []string{}
	}
	if total < 2 {
		return []string{"And all for the want of a nail."}
	}
	phrase := make([]string, total)
	for i := 0; i < total-1; i++ {
		phrase[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
	}
	phrase[total-1] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	return phrase
}
