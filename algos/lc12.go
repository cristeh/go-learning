package algos

import "strings"

func intToRoman(num int) string {
	romanToInt := []struct {
		Number  int
		Mapping string
	}{
		{
			Number:  1000,
			Mapping: "M",
		},
		{
			Number:  900,
			Mapping: "CM",
		},
		{
			Number:  500,
			Mapping: "D",
		},
		{
			Number:  400,
			Mapping: "CD",
		},
		{
			Number:  100,
			Mapping: "C",
		},
		{
			Number:  90,
			Mapping: "XC",
		},
		{
			Number:  50,
			Mapping: "L",
		},
		{
			Number:  40,
			Mapping: "XL",
		},
		{
			Number:  10,
			Mapping: "X",
		},
		{
			Number:  9,
			Mapping: "IX",
		},
		{
			Number:  5,
			Mapping: "V",
		},
		{
			Number:  4,
			Mapping: "IV",
		},
		{
			Number:  1,
			Mapping: "I",
		},
	}
	var sb strings.Builder
	for _, mapping := range romanToInt {
		for num >= mapping.Number {
			sb.WriteString(mapping.Mapping)
			num -= mapping.Number
		}
	}
	return sb.String()
}
