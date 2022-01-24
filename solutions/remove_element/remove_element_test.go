package solutions

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	t.Run("Remove duplicates from sorted array", func(t *testing.T) {
		nums := []int{3,2,2,3}
		val := 3

		expectedArray := []int{2,2}

		k := RemoveElement(nums, val)
		if k != len(expectedArray) {
			t.Errorf("Expected to receive %d, but got: %d", len(expectedArray), k)
		}

		for i, v := range expectedArray {
			if v != nums[i] {
				t.Errorf("Expected to nums[%d] to be %d, but %d received", i, v, nums[i])
			}
		}
	})

	t.Run("Remove duplicates from sorted array", func(t *testing.T) {
		nums := []int{0,1,2,2,3,0,4,2}
		val := 2

		expectedArray := []int{0, 1, 3, 0, 4}

		k := RemoveElement(nums, val)
		if k != len(expectedArray) {
			t.Errorf("Expected to receive %d, but got: %d", len(expectedArray), k)
		}

		for i, v := range expectedArray {
			if v != nums[i] {
				t.Errorf("Expected to nums[%d] to be %d, but %d received", i, v, nums[i])
			}
		}
	})
}
