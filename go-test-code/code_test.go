package go_test_code

import (
	//自定义包名aodeng
	aodeng "fmt"
	"math"
	"os"
	"runtime"
	//test包
	"testing"
)

//类型转换
func TestCode(t *testing.T) {
	a := 5.0
	//类型 B 的值 = 类型 B(类型 A 的值)
	b := int(a)
	aodeng.Print(b)
}

//常量
func TestCode2(t *testing.T) {
	const Pi = 3.14159
	//显式类型定义：
	const b1 string = "abc"
	//隐式类型定义：
	const b2 = "abc"
	t.Log(Pi)
	t.Log(b1)
	t.Log(b2)
}

//全局变量
var (
	a int
	b bool
	c string
)

//下面是如何通过runtime包在运行时获取所在的操作系统类型，以及如何通过 os 包中的函数 os.Getenv() 来获取环境变量中的值，并保存到 string 类型的局部变量 path 中
func TestCode3(t *testing.T) {
	var goos string = runtime.GOOS
	aodeng.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	aodeng.Printf("Path is %s\n", path)
}

//变量除了可以在全局声明中初始化，也可以在 init 函数中初始化。这是一类非常特殊的函数，它不能够被人为调用，而是在每个包完成初始化后自动执行，并且执行优先级比 main 函数高。
//每个源文件都只能包含一个 init 函数。初始化总是以单线程执行，并且按照包的依赖关系顺序执行。
//一个可能的用途是在开始执行程序之前对数据进行检验或修复，以保证程序状态的正确性
var Pi float64

func init() {
	Pi = 4 * math.Atan(1)
}
func TestCode4(t *testing.T) {
	var twoPi = 2 * Pi
	aodeng.Print("2*Pi:", twoPi)
}
