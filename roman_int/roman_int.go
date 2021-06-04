package roman_int

import (
	"fmt"
)

// RomanToInt 罗马数字转整数
func RomanToInt(s string) int {
	var sum int
	preNum := getValue(fmt.Sprintf("%c", s[0]))
	for k, v := range s {
		if k == 0 {
			continue
		}
		num := getValue(fmt.Sprintf("%c", v))
		if preNum < num {
			sum -= preNum
		} else {
			sum += preNum
		}
		preNum = num
	}
	sum += preNum
	return sum
}

func getValue(s string) int {
	switch s {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0
	}
}
