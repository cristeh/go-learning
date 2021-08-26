package algos

func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	left[0] = 1
	right[len(nums)-1] = 1
	for start, end := 1, len(nums)-2; start < len(nums); start, end = start+1, end-1 {
		left[start] = left[start-1] * nums[start-1]
		right[end] = right[end+1] * nums[end+1]
	}
	result := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		result = append(result, left[i]*right[i])
	}
	return result
}
