package add_two_nums

import (
	"github.com/stretchr/testify/assert"
	"log"
	"math"
	"testing"
)

func TestPow(t *testing.T) {

	x := 2
	expected := 100
	actual := Pow10(x)
	assert.Equal(t, expected, actual)

	x = 4
	expected = int(math.Pow10(x))
	actual = Pow10(x)
	assert.Equal(t, expected, actual)

}

func TestAddTwoNumbers(t *testing.T) {
	l111 := &ListNode{
		Val:  3,
		Next: nil,
	}
	l11 := &ListNode{
		Val:  4,
		Next: l111,
	}
	l1 := &ListNode{
		Val:  2,
		Next: l11,
	}
	l222 := &ListNode{
		Val:  4,
		Next: nil,
	}
	l22 := &ListNode{
		Val:  6,
		Next: l222,
	}
	l2 := &ListNode{
		Val:  5,
		Next: l22,
	}

	sum := AddTwoNumbers(l1, l2)
	log.Println(sum)

}
