package is_palindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	x := 101
	expected := true
	actual := IsPalindrome(x)
	assert.Equal(t, expected, actual)

	x = -101
	expected = false
	actual = IsPalindrome(x)
	assert.Equal(t, expected, actual)
}
