package int_reverse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {

	x := 123
	expected := 321
	actual := Reverse(x)
	assert.Equal(t, expected, actual)

	x = 210
	expected = 12
	actual = Reverse(x)
	assert.Equal(t, expected, actual)

}
