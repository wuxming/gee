package min

import (
	"testing"
)

func BenchmarkMin(b *testing.B) {
	m := Default()
	m.GET("/benchmark/test", func(c *Context) {
		c.Params("223")
		c.JSON(200, H{})
	})
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			go func() {
				PreformRequset(m, "GET", "/benchmark/test")
			}()
		}
	}
	b.StopTimer()
	b.ReportAllocs()
}

// go test -run="benchmark_test.go" -test.bench="BenchmarkMin"
// 				执行次数   平均时间		均内存大小分配 平均内存大小次数  执行时间
// 				24169    56867 ns/op    6149 B/op    37 allocs/op    2.215s
