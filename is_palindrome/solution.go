package is_palindrome

import "strconv"

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)

	r := []rune(s)

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	y, _ := strconv.Atoi(string(r))

	if x == y {
		return true
	} else {
		return false
	}

}

/*还有第二种，数组前后位置比较*/
