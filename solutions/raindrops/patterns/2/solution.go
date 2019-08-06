// Package raindrops is used to transform a number into the lovely sound of rain
package raindrops

import (
	"strconv"
)

var m = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Convert converts a number to rain drop speech.
func Convert(num int) string {
	var res string
	for i := 3; i < 8; i += 2 {
		if num%i == 0 {
			res += m[i]
		}
	}

	if res == "" {
		return strconv.Itoa(num)
	}

	return res
}
