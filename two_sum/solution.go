package two_sum

// TwoSum 两数之和
func TwoSum(nums []int, target int) []int {

	idNumMap := make(map[int]int)

	for idx, num := range nums {
		needNum := target - num

		idn, ok := idNumMap[needNum]
		if ok {
			return []int{idn, idx}
		}

		idNumMap[num] = idx
	}

	return []int{}

}
