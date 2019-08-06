// Package hamming provides the Distance function.
package hamming

import "fmt"

// Distance calculates the Hamming distance between two DNA strands.
func Distance(a, b string) (int, error) {
	var (
		aRunes = []rune(a)
		bRunes = []rune(b)
	)

	if len(aRunes) != len(bRunes) {
		return 0, fmt.Errorf("DNA strands %s and %s have different length", a, b)
	}

	diff := 0
	for i, r := range aRunes {
		if r != bRunes[i] {
			diff++
		}
	}
	return diff, nil
}
