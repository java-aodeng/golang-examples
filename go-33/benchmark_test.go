package go_33

import (
	"bytes"
	"testing"
)

/*
Benchmark
用于一些性能测试中，方法名字以Benchmark为前缀
参数使用testing.B

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	//与性能测试无关代码
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//测试代码
	}
	b.StopTimer()
	//与性能测试无关代码
}

*/

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer

		for _, elem := range elems {
			buf.WriteString(elem)

		}
	}
	b.StopTimer()

}

/*
//运行结果
goos: windows
goarch: amd64
BenchmarkConcatStringByBytesBuffer-4   	10000000	       142 ns/op
PASS*/
