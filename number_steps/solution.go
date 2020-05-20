package number_steps

func NumberOfSteps(num int) int {
	i := 0

	for num != 0 {
		i += 1
		if num%2 == 0 {
			num = num / 2
		} else {
			num = num - 1
		}

	}
	return i
}
