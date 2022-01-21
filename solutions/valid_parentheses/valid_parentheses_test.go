package solutions

import "testing"

func TestValidParentheses(t *testing.T)  {
	t.Run("Valid Parentheses", func(t *testing.T) {
		s := "[({})]"
		if IsValid(s) == false {
			t.Errorf("Expected to be true for value: %s", s)
		}
	})

	t.Run("Valid Parentheses", func(t *testing.T) {
		s := "[({)})]"
		if IsValid(s) == true {
			t.Errorf("Expected to be false for value: %s", s)
		}
	})

	t.Run("Valid Parentheses", func(t *testing.T) {
		s := ""
		if IsValid(s) == false {
			t.Errorf("Expected to be true for value: %s", s)
		}
	})

	t.Run("Valid Parentheses", func(t *testing.T) {
		s := "()"
		if IsValid(s) == false {
			t.Errorf("Expected to be true for value: %s", s)
		}
	})
	t.Run("Valid Parentheses", func(t *testing.T) {
		s := "]["
		if IsValid(s) == true {
			t.Errorf("Expected to be false for value: %s", s)
		}
	})
}