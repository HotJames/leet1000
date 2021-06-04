package two_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	// GiB // 1073741824
	// TiB // 1099511627776             (超过了int32的范围)
	// PiB // 1125899906842624
	// EiB // 1152921504606846976
	// ZiB // 1180591620717411303424    (超过了int64的范围)
	// YiB // 1208925819614629174706176
)

func TestTwoSum(t *testing.T) {
	nums := []int{3, 3}
	target := 6
	expected := []int{0, 1}

	actual := TwoSum(nums, target)

	assert.Equal(t, expected, actual)

}

func BenchmarkTwoSum(b *testing.B) {
	//b.ReportAllocs()
	b.ResetTimer()
	nums := []int{3, 3}
	target := 6

	TwoSum(nums, target)
}
