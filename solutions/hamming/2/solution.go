package hamming

import (
	"errors"
	"strings"
)

// Distance calculates the hamming distance
func Distance(a, b string) (int, error) {
	cnt := 0
	a = strings.TrimSpace(a)
	b = strings.TrimSpace(b)

	if len(a) != len(b) {
		return 0, errors.New("lengths of inputs are not the same")
	}
	for k, v := range a {
		if string(v) != string(b[k]) {
			cnt++
		}

	}
	return cnt, nil
}
