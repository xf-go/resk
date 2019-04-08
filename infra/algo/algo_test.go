package algo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSimpleRand(t *testing.T) {
	ForTest("简单随机算法", t, SimpleRand)
}

func TestBeforeShuffle(t *testing.T) {
	ForTest("先洗牌算法", t, BeforeShuffle)
}

func TestDoubleRandom(t *testing.T) {
	ForTest("二次随机算法", t, DoubleRandom)
}

func TestDoubleAverage(t *testing.T) {
	ForTest("二倍均值算法", t, DoubleAverage)
}

func ForTest(message string, t *testing.T, fn func(count, amount int64) int64) {
	count, amount := int64(10), int64(10000)
	remain := amount
	sum := int64(0)
	for i := int64(0); i < count; i++ {
		x := fn(count-i, remain)
		remain -= x
		sum += x
	}
	Convey(message, t, func() {
		So(sum, ShouldEqual, amount)
	})
}
