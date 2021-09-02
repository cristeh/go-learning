package algos

func exist(board [][]byte, word string) bool {
	for i, _ := range board {
		for j, _ := range board[i] {
			if backtracking(i, j, board, 0, word) {
				return true
			}
		}
	}
	return false
}

func backtracking(row int, col int, board [][]byte, wordIndex int, word string) bool {
	if wordIndex == len(word) {
		return true //We backtracked up to a successful word, solution is true
	}
	//if we step out of bounds or charAtPos doesn't match the one in word, this is not the solution
	if row >= len(board) || row < 0 || col >= len(board[0]) || col < 0 || word[wordIndex] != board[row][col] {
		return false
	}
	//Visit clockwise
	offset := []struct {
		RowOffset int
		ColOffset int
	}{
		{0, 1},  //Right
		{1, 0},  // Down
		{0, -1}, //Left
		{-1, 0}, //Up
	}
	currChar := board[row][col]
	board[row][col] = '$' //Mark visited while attempting path
	for _, pos := range offset {
		if backtracking(row+pos.RowOffset, col+pos.ColOffset, board, wordIndex+1, word) {
			return true
		}
	}
	//Backtrack
	board[row][col] = currChar
	return false
}
