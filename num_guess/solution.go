package num_guess

func Game(guess []int, answer []int) int {
	x := 0

	for i := 0; i < 4; i++ {
		if guess[i] == answer[i] {
			x++
		}
	}
	return x
}
