package algos

func twoSumII(numbers []int, target int) []int {
	start, end := 0, len(numbers)-1
	for start < end {
		currSum := numbers[start] + numbers[end]
		if currSum == target {
			return []int{start + 1, end + 1}
		}
		if currSum < target {
			start++
		}
		if currSum > target {
			end--
		}
	}
	return []int{-1, -1}
}
