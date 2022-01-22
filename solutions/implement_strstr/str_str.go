package solutions

func StrStr(haystack string, needle string) int {
	needleLen := len(needle)
	haystackLen := len(haystack)
	if needleLen == 0 {
		return 0
	}

	for i := 0; i < haystackLen; i++ {
		if haystackLen < i + needleLen {
			return -1
		}
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}
	return -1
}