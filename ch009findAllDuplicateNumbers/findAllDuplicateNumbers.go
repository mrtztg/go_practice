package findallduplicatenumbers

func findallduplicatenumbers(nums []int) []int {
	for i := 0; i < len(nums); {
		if nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}
	duplicates := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if i+1 != nums[i]{
			duplicates = append(duplicates, nums[i])
		}
	}
	return duplicates
}
