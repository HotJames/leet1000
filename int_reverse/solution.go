package int_reverse

import (
	"math"
)

// Reverse 整数反转
func Reverse(x int) int {

	sign := 1

	if x < 0 {
		sign = -1
		x = sign * x
	}

	y := 0

	for x > 0 {

		temp := x % 10

		y = y*10 + temp

		x = x / 10

	}

	y = sign * y

	if y > math.MaxInt32 || y < math.MinInt32 {
		return 0
	}

	return y
}
