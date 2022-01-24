package solutions

import "strings"

func romanToInt(s string) int {
	romanDigits := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	subDigits := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	result := 0
	strlen := len(s)
	digits := strings.Split(s, "")

	for i := 0; i < strlen; i++ {
		key := ""
		nextIdx := i + 1

		if nextIdx < strlen {
			key = digits[i] + digits[nextIdx]
		}

		if val, ok := subDigits[key]; ok {
			result += val
			i++
		} else {
			result += romanDigits[digits[i]]
		}
	}
	return result
}
