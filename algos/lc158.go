package algos

func lengthOfLongestSubstringTwoDistinct(s string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	start := 0
	end := 0
	freqMap := make(map[byte]int)
	result := 0
	for end < len(s) {
		endChar := s[end]
		freqMap[endChar]++
		for len(freqMap) > 2 {
			startChar := s[start]
			freqMap[startChar]--
			if freqMap[startChar] == 0 {
				delete(freqMap, startChar)
			}
			start++
		}
		result = max(result, end-start+1)
		end++
	}
	return result
}
