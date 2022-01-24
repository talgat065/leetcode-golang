package solutions

func RemoveDuplicates(nums []int) int {
	var sorted, lastIndex int
	if len(nums) > 0 {
		sorted++
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[lastIndex] {
			lastIndex++; sorted++
			nums[lastIndex] = nums[i]
		}
	}
	return sorted
}