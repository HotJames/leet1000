package remove_duplicates

import (
	"strconv"
	"strings"
)

func RemoveDuplicates(nums []int) int {
	s := strconv.Itoa(nums[0])
	l := len(nums)
	for i := 1; i < l; i++ {
		j := strconv.Itoa(nums[i])
		if !strings.Contains(s, j) {
			s += j
		} else {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(strings.Split(s, ""))
}
