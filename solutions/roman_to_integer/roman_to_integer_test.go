package solutions

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	t.Run("Remove duplicates from sorted array", func(t *testing.T) {
		i := romanToInt("III")
		answer := 3
		if i != answer {
			t.Errorf("Expected to receive %d, but got: %d", answer, i)
		}
	})
	t.Run("Remove duplicates from sorted array", func(t *testing.T) {
		i := romanToInt("IV")
		answer := 4
		if i != answer {
			t.Errorf("Expected to receive %d, but got: %d", answer, i)
		}
	})
	t.Run("Remove duplicates from sorted array", func(t *testing.T) {
		i := romanToInt("LVIII")
		answer := 58
		if i != answer {
			t.Errorf("Expected to receive %d, but got: %d", answer, i)
		}
	})
	t.Run("Remove duplicates from sorted array", func(t *testing.T) {
		i := romanToInt("MCMXCIV")
		answer := 1994
		if i != answer {
			t.Errorf("Expected to receive %d, but got: %d", answer, i)
		}
	})
}
