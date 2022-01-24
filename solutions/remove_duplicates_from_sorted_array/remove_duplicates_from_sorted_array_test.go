package solutions

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	t.Run("Remove duplicates from sorted array", func(t *testing.T) {
		nums := []int{1, 1, 2}
		k := RemoveDuplicates(nums)
		if k != 2 {
			t.Errorf("Expected to receive %d, but got: %d", 2, k)
		}

		expectedArray := []int{1,2}
		for i, v := range expectedArray {
			if v != nums[i] {
				t.Errorf("Wrong numbers order")
			}
		}
	})

	t.Run("Remove duplicates from sorted array", func(t *testing.T) {
		nums := make([]int, 0, 0)
		k := RemoveDuplicates(nums)
		if k != 0 {
			t.Errorf("Expected to receive %d, but got: %d", 0, k)
		}
	})

	t.Run("Remove duplicates from sorted array", func(t *testing.T) {
		nums := []int{1}
		k := RemoveDuplicates(nums)
		if k != 1 {
			t.Errorf("Expected to receive %d, but got: %d", 1, k)
		}
		if nums[0] != 1 {
			t.Errorf("Wrong numbers order")
		}
	})

	t.Run("Remove duplicates from sorted array", func(t *testing.T) {
		nums := []int{1, 1}
		k := RemoveDuplicates(nums)
		if k != 1 {
			t.Errorf("Expected to receive %d, but got: %d", 1, k)
		}
		expectedArray := []int{1}
		for i, v := range expectedArray {
			if v != nums[i] {
				t.Errorf("Wrong numbers order")
			}
		}
	})

	t.Run("Remove duplicates from sorted array", func(t *testing.T) {
		nums := []int{0,0,1,1,1,2,2,3,3,4}
		k := RemoveDuplicates(nums)
		if k != 5 {
			t.Errorf("Expected to receive %d, but got: %d", 5, k)
		}
		expectedArray := []int{0,1,2,3,4}
		for i, v := range expectedArray {
			if v != nums[i] {
				t.Errorf("Wrong numbers order")
			}
		}
	})
}
