package ch008findduplicatenumber

func findDuplicateNumber(nums []int) int{
	for i := 0; i < len(nums); {
		if nums[i] != nums[nums[i] - 1] {
			nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
		} else if i + 1 != nums[i] {
			return nums[i]
		} else {
			i++
		}
	}
	return -1
}