package algo

import (
	"math/rand"
	"time"
)

// 二倍均值算法
func DoubleAverage(count, amount int64) int64 {
	if count == 1 {
		return amount
	}
	// 计算出最大可用金额
	max := amount - min*count
	// 计算最大可用平均值
	avg := max / count
	// 二倍均值基础上加上最小金额,防止出现0值
	avg2 := 2*avg + min
	// 随机红包金额序列元素,把二倍均值作为随机的最大数
	rand.Seed(time.Now().UnixNano())
	x := rand.Int63n(avg2) + min
	return x
}
