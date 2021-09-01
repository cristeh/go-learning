package algos

func generateParenthesis(n int) []string {
	var result []string
	var backtracking func(currentIteration string, open int, closed int, n int)
	backtracking = func(currentIteration string, open int, closed int, n int) {
		if open+closed == n*2 {
			result = append(result, currentIteration)
			return
		}
		if open < n {
			currentIteration += "("
			backtracking(currentIteration, open+1, closed, n)
			currentIteration = currentIteration[:len(currentIteration)-1]
		}
		//Close current open iterations
		if closed < open {
			currentIteration += ")"
			backtracking(currentIteration, open, closed+1, n)
			currentIteration = currentIteration[:len(currentIteration)-1]
		}
	}
	backtracking("", 0, 0, n)
	return result
}

/*
Tree expansion for 3: add (, backtrack, add ), try add ( again and repeat
( ( ( ) ) )
( ( ) ( ) )
...
( ) ( ) ( )
*/
