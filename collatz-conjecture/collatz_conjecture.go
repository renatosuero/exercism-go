package collatzconjecture

import (
	"errors"
)

// CollatzConjecture return the number of steps required to reach 1.
func CollatzConjecture(value int) (int, error) {
	if value <= 0 {
		return value, errors.New("Invalid number")
	}
	steps := 0
	for value != 1 {
		value = processNumber(value)
		steps++
	}
	return steps, nil
}
func processNumber(val int) int {
	if val%2 == 0 {
		return val / 2
	}
	return (val * 3) + 1
}
