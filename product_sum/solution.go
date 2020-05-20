package product_sum

func SubtractProductAndSum(n int) int {

	pro := 1
	sub := 0

	for n > 0 {

		temp := n % 10

		pro = pro * temp

		sub = sub + temp

		n = n / 10
	}
	return pro - sub
}
