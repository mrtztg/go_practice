package ch010findcorruptpair

func findCorruptPair(nums []int) []int {
	for i := 0; i < len(nums); {
		if nums[i] != nums[nums[i] - 1] {
			nums[i], nums[nums[i] -1 ] = nums[nums[i] -1], nums[i]
		} else {
			i++
		}
	}
	corruptions := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if i + 1 != nums[i] {
			corruptions = append(corruptions, nums[i], i+1)
		}
	}
	return corruptions
}