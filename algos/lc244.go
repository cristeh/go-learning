package algos

type WordDistance struct {
	positions map[string][]int
}

func Lc244Constructor(wordsDict []string) WordDistance {
	positions := make(map[string][]int)
	for i, word := range wordsDict {
		positions[word] = append(positions[word], i)
	}
	return WordDistance{
		positions: positions,
	}
}

func (this *WordDistance) Shortest(word1 string, word2 string) int {
	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	abs := func(a int) int {
		if a < 0 {
			return -a
		} else {
			return a
		}
	}
	w1Pointer, w2Pointer := 0, 0
	minDistance := 300001
	for w1Pointer < len(this.positions[word1]) && w2Pointer < len(this.positions[word2]) {
		minDistance = min(minDistance, abs(this.positions[word1][w1Pointer]-this.positions[word2][w2Pointer]))
		if this.positions[word1][w1Pointer] < this.positions[word2][w2Pointer] {
			w1Pointer++
		} else {
			w2Pointer++
		}
	}
	return minDistance
}

/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(wordsDict);
 * param_1 := obj.Shortest(word1,word2);
 */

//practice: 0
//makes: 1 4
//perfect: 2
//coding: 3

//coding, practice

//3 6 8 12
//2 7 9
