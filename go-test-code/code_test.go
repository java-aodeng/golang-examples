package go_test_code

import (
	aodeng "fmt"
	"testing"
)

//类型转换
func TestCode(t *testing.T) {
	a := 5.0
	//类型 B 的值 = 类型 B(类型 A 的值)
	b := int(a)
	aodeng.Print(b)
}
