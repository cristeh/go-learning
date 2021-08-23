package algos

func twoSum(nums []int, target int) []int {
	revIdx := make(map[int]int)
	for i, ele := range nums {
		if j, ok := revIdx[target-ele]; ok {
			return []int{i, j}
		}
		revIdx[ele] = i
	}
	return []int{-1, -1}
}
