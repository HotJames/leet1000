package add_two_nums

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	i := 0
	sum1 := 0
	sum2 := 0
	for true {
		sum1 += l1.Val * Pow10(i)
		if l1.Next != nil {
			l1 = l1.Next
		} else {
			break
		}
		i += 1
	}
	i = 0
	for true {
		sum2 += l2.Val * Pow10(i)
		if l2.Next != nil {
			l2 = l2.Next
		} else {
			break
		}
		i += 1
	}

	sum := sum1 + sum2

}

func Pow10(n int) (s int) {
	s = 1
	for i := 0; i < n; i++ {
		s = s * 10
	}
	return
}
