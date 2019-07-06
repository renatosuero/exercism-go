package grains

import "fmt"

//Square how many grains where on each square
func Square(value int) (uint64, error) {
	if value < 1 || value > 64 {
		return 0, fmt.Errorf("Value needs to be between 1 and 64, the value received was %d", value)
	}
	return 1 << uint(value-1), nil
}

//Total returns total number of grains
func Total() uint64 {
	return 1<<64 - 1
}
