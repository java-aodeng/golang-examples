package go_02

import "testing"

//定义常量

//快速设置连续值,定义一周七天，值连续+1
const(
	Monday =iota +1
	Tuesday
	Wedensday
	Thurday
	Friday
	Saturday
	Sunday
)

//第一个比特位为1 其他位为0 表示Open 依此类推...
const (
	Open =1 << iota
	Close
	Pending
)

//输出变量值 输出为1，6
func TestConstantTry(t *testing.T)  {
	t.Log(Monday)
	t.Log(Saturday)
}

//比特位比较 1的二进制为0001 输出 true false false
func TestConstantTry1(t *testing.T)  {
	a:=1
	t.Log(a&Open==Open,a&Close==Close,a&Pending==Pending)
}