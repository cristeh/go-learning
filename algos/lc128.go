package algos

func longestConsecutive(nums []int) int {
	contains := func(e int, m map[int]bool) bool {
		_, ok := m[e]
		return ok
	}
	elements := make(map[int]bool)
	for _, num := range nums {
		elements[num] = true
	}
	var maxSequence = 0
	for num, _ := range elements {
		localMaxSequence := 1
		for prev, ok := num-1, contains(num-1, elements); ok; prev, ok = prev-1, contains(prev-1, elements) {
			delete(elements, prev)
			localMaxSequence++
		}
		for next, ok := num+1, contains(num+1, elements); ok; next, ok = next+1, contains(next+1, elements) {
			delete(elements, next)
			localMaxSequence++
		}
		if maxSequence < localMaxSequence {
			maxSequence = localMaxSequence
		}
	}
	return maxSequence
}
