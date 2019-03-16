package hamming

import "errors"

// Distance returns the Hamming distance of 2 strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strings are different lengths")
	}

	var diff int
	for i := range a {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff, nil
}
