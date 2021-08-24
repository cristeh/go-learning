package algos

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {
	num1Parts := strings.Split(num1, "+")
	num2Parts := strings.Split(num2, "+")
	//Let's assume (a + bi) * (c + di) =  (ac - bd) + (ad + cb)i
	a, _ := strconv.Atoi(num1Parts[0])
	b, _ := strconv.Atoi(num1Parts[1][:len(num1Parts[1])-1])
	c, _ := strconv.Atoi(num2Parts[0])
	d, _ := strconv.Atoi(num2Parts[1][:len(num2Parts[1])-1])
	return fmt.Sprint(a*c-b*d, "+", a*d+c*b, "i")
}
