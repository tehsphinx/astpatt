// Package hamming provides the Distance function.
package hamming

import "fmt"

// Distance calculates the Hamming distance between two DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("DNA strands %s and %s have different length", a, b)
	}

	diff := 0
	bRunes := []rune(b)
	for i, r := range a {
		if r != bRunes[i] {
			diff++
		}
	}
	return diff, nil
}
