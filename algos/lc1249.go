package algos

import "strings"

type IndexedParantheses struct {
	Symbol rune
	Index  int
}

func minRemoveToMakeValid(s string) string {
	stack := make([]IndexedParantheses, 0, len(s))
	for i, char := range s {
		if char != '(' && char != ')' {
			continue
		}
		if char == ')' && len(stack) > 0 && stack[len(stack)-1].Symbol == '(' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, IndexedParantheses{
				Symbol: char,
				Index:  i,
			})
		}
	}
	var sb strings.Builder
	for i, ch := range s {
		if len(stack) > 0 && stack[0].Index == i {
			stack = stack[1:]
		} else {
			sb.WriteRune(ch)
		}
	}
	return sb.String()
}
