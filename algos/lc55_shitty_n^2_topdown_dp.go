package algos

const (
	UnknownValue = iota
	TrueValue
	FalseValue
)

func canJump(nums []int) bool {
	memo := make([]int, len(nums))
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	var backtrackingCanJump func(position int, nums []int) bool
	backtrackingCanJump = func(position int, nums []int) bool {
		if memo[position] == TrueValue {
			return true
		}
		if memo[position] == FalseValue {
			return false
		}
		if position == len(nums)-1 {
			return true
		}
		finalPosition := min(position+nums[position], len(nums)-1)
		for curr := position + 1; curr <= finalPosition; curr++ {
			step := backtrackingCanJump(curr, nums)
			if step {
				memo[curr] = TrueValue
				return true
			} else {
				memo[curr] = FalseValue
			}
		}
		return false
	}
	return backtrackingCanJump(0, nums)
}
