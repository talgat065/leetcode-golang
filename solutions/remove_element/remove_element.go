package solutions

// RemoveElement nums = [3,2,2,3], val = 3
// RemoveElement nums = [0,1,2,2,3,0,4,2], val = 2
func RemoveElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}