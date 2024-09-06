package ch006findMissingNumber

type Solution struct {
}

func (s *Solution) findMissingNumber(nums []int) int {
	i := 0
	for i < len(nums) {
		if i != nums[i] && nums[i] != len(nums) {
			tempNumber := nums[i]
			nums[i] = nums[tempNumber]
			nums[tempNumber] = tempNumber
		} else {
			i++
		}
	}
	for i = 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}
