package ch007findAllMissingNumbers

type solution struct {
}

func (s *solution) findAllMissingNumbers(nums []int) []int {
	for i := 0; i < len(nums); {
		if nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}
	missingNumbers := make([]int, 0)
	for i, number := range nums {
		if i+1 != number {
			missingNumbers = append(missingNumbers, i+1)
		}
	}
	return missingNumbers
}

// 2 3 1 8 2 3 5 1
// 3 2 1 8 2 3 5 1
// 1 2 3 8 2 3 5 1
// 1 2 3 1 2 3 5 8
// 1 2 3 1 5 3 2 8
