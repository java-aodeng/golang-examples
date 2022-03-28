package go_13

import (
	"testing"
)

//定义一个接口
type Programmer interface {
	WriteHelloWorld() string
}

//实现接口
type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "hello world"
}

func Test(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
