package roman_int

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	s := "LVIII"
	expected := 58
	actual := RomanToInt(s)
	assert.Equal(t, expected, actual)
}
