package hamming

import (
	"errors"
)

// Distance Calculate the Hamming difference between two DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strings have different size")
	}
	totalDistance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			totalDistance++
		}
	}
	return totalDistance, nil
}
