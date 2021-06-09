package plus_one

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlusOne(t *testing.T) {

	digits := []int{1, 2, 3}
	expected := []int{1, 2, 4}
	actual := PlusOne(digits)
	assert.Equal(t, expected, actual)

	digits = []int{9, 9, 9}
	expected = []int{1, 0, 0, 0}
	actual = PlusOne(digits)
	assert.Equal(t, expected, actual)

}
