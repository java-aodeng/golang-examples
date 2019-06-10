package go_02

import (
	"testing"
)

//定义变量

//go语言实现一个Fibonacci数列
func TestFib(t *testing.T) {
	a:=1
	b:=1
	t.Log(a)
	for i:=0;i<5 ;i++  {
		t.Log("",b)
		temp:=a
		a=b
		b=temp+a
	}
}

//go语言多变量赋值 ，一个赋值语句中实现对多个变量赋值
func TestExchange(t *testing.T)  {
	a:=1
	b:=2
	a,b=b,a
	t.Log(a,b)
}