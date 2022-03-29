package go_14

import (
	"fmt"
	"testing"
)

//go语言扩展类似java的继承 但又不完全是
//这是一只猪类
type Pig struct {
}

func (p *Pig) Speak() {
	fmt.Print("我是猪")
}

func (p *Pig) MyNameIs(name string) {
	p.Speak()
	fmt.Print("my name is" + name)
}

//这是一条狗类 继承于猪类
type Dog struct {
	Pig
}

func (d *Dog) Speak() {
	//类似于java的重载
	fmt.Print("我不是猪，我是狗")
}

func (d *Dog) MyNameIs(name string) {
	d.Speak()
	fmt.Print("my name is" + name)
}

//这是一个傻逼
type SB struct {
	Pig
}

func (s *SB) Speak() {
	fmt.Print("我是一个傻逼")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.MyNameIs("你爸爸")

	fmt.Print("========\n")

	sb := new(SB)
	sb.MyNameIs("你爷爷")
}

//运行结果
//=== RUN   TestDog
//我不是猪，我是狗my name is你爸爸========
//我是猪my name is你爷爷--- PASS: TestDog (0.00s)
//PASS
