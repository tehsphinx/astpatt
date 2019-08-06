// Package hamming contains function to calculate amount of character difference between two equal strings.
package hamming

// NotEqualLengthError is raised if strings supplied to Distance function have different length.
type NotEqualLengthError struct{}

func (error *NotEqualLengthError) Error() string {
	return "strings have different length"
}

// Distance calculates amount of character difference between two equal strings.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, &NotEqualLengthError{}
	}

	var diff int
	for i := range a {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff, nil

}
