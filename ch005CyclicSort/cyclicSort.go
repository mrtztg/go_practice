package ch005CyclicSort

type Solution struct{}

func (s *Solution) sort(Nums []int) []int {
	if len(Nums) < 2 {
		return Nums
	}
	for i := 0; i < len(Nums); i++ {
		for Nums[i] != i+1 {
			tempNo := Nums[i]
			Nums[i] = Nums[tempNo-1]
			Nums[tempNo-1] = tempNo
		}
	}
	return Nums
}
