// Package raindrops is used to transform a number into the lovely sound of rain
package raindrops

import "strconv"

// Convert converts a number to rain drop speech.
func Convert(num int) string {
	var res string
	for _, rule := range speechDef {
		if num%rule.val == 0 {
			res += rule.speech
		}
	}

	if res == "" {
		return strconv.Itoa(num)
	}

	return res
}

var speechDef = []struct {
	val    int
	speech string
}{
	{val: 3, speech: "Pling"},
	{val: 5, speech: "Plang"},
	{val: 7, speech: "Plong"},
}
