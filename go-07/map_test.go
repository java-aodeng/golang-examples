package go_07

import (
	aodeng "fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := map[string]int{"a": 1, "aa": 2, "aaa": 3}
	t.Log(m["a"])
}

func TestTraveMap(t *testing.T) {
	m := map[int]int{1: 1, 2: 4, 3: 9}

	for k, v := range m {
		t.Log(k, v)
	}
}

func Test1(a *testing.T) {
	for i := 0; i < 500; i++ {
		//aa:=1
		//bb:=2
	}
	//根据你的个人喜好对包名进行重新设置
	aodeng.Print("this is aodeng pkg")
	aodeng.Print("\n")
	aodeng.Print(functionName(2))
	print("aaa")
	main()
}

//这是定义一个函数最简单的格式
func Sum(a, b int) int {
	return a + b
}

//因此符合规范的函数一般写成如下的形式
func functionName(a int) int {
	return Sum(a, a)
}

func main() {
	//如何在 Go 中使用国际化字符
	aodeng.Printf("Καλημέρα κόσμε; or こんにちは 世界\n")
}
