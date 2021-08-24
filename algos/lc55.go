package algos

import "sort"

type ByIntervalStart [][]int

func (a ByIntervalStart) Len() int           { return len(a) }
func (a ByIntervalStart) Less(i, j int) bool { return a[i][0] < a[j][0] }
func (a ByIntervalStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func merge(intervals [][]int) [][]int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	sort.Sort(ByIntervalStart(intervals))
	result := make([][]int, 0, len(intervals))
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		lastResultInterval := result[len(result)-1]
		if lastResultInterval[1] >= intervals[i][0] {
			lastResultInterval[1] = max(intervals[i][1], lastResultInterval[1])
		} else {
			result = append(result, intervals[i])
		}
	}
	return result
}
