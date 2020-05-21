package min_count

func MinCount(coins []int) int {
	i := 0
	for _, x := range coins {
		if x%2 == 0 {
			i += x / 2
		} else {
			i = i + x/2 + 1
		}
	}
	return i
}
