package go_15

import (
	"fmt"
	"testing"
)

//go语言的空接口与断言
//1.go语言的空接口可以表示任何类型，
//2.通过断言来将空接口转换为制定类型
//如：v,ok :=p.(int) //ok=true时为转换成功
func DoSomething(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Print("Integer类型的", v)
	case string:
		fmt.Print("String类型的", v)
	default:
		fmt.Print("未知类型的")
	}
}

//这里主要是体会一下go语言的多态 与其他语言的区别
func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(1)
	fmt.Print("\n")
	DoSomething("1")
	fmt.Print("\n")
	DoSomething(1.1)
}

//运行结果
//=== RUN   TestEmptyInterfaceAssertion
//Integer类型的1
//String类型的1
//未知类型的--- PASS: TestEmptyInterfaceAssertion (0.     00s)
//PASS
