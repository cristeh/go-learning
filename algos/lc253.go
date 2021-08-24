package algos

import "sort"

func minMeetingRooms(intervals [][]int) int {
	startTimes, endTimes := make([]int, 0, len(intervals)), make([]int, 0, len(intervals))
	for _, interval := range intervals {
		startTimes = append(startTimes, interval[0])
		endTimes = append(endTimes, interval[1])
	}
	sort.Ints(startTimes)
	sort.Ints(endTimes)
	result := 0
	for start, end := 0, 0; start < len(intervals); start++ {
		if startTimes[start] < endTimes[end] {
			result++
		} else {
			end++
		}
	}
	return result
}

//0 5 15
//10 20 30
