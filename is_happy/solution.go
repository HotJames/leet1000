package is_happy

func IsHappy(n int) bool {

	site := 1

	for a := 1; a*10 <= n; a = a * 10 {
		site += 1
	}

	return false
}
