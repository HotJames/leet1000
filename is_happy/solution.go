package is_happy

import (
	"fmt"
)

func IsHappy(n int) bool {

	site := 1

	for a := 1; a*10 <= n; a = a * 10 {
		site += 1
	}

	fmt.Println(site)
	return false
}
