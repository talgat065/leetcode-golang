package solutions

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}

	result := LongestCommonPrefix(strs)
	if result != "fl" {
		t.Errorf("Expected to be: %s, but received: %s instead", "fl", result)
	}
}
