package go_16

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

//Go的错误机制
//1、这里需要注意：go没有异常机制，都自己写判断实现的
//2、error类型实现error接口 type error interface{Error() string}
//3、可以通过errors.New来创建错误实例

//这里创建两个错误实列，这一步，类似于java的自定义异常
var LessThanTwoError = errors.New("参数小于 2")
var LargerThenHundredError = errors.New("参数大于 100")

//定义一个斐波那契列方法，这个方法前面做逻辑判断，返回自定义error实列
func GetFibonacci(n int) ([]int, error) {
	//参数校验，不满足返回定义的自定义异常
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThenHundredError
	}
	fibList := []int{1, 1}
	//实现个小学数学题，斐波那契列方法
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

//优化处理，及早失败，避免嵌套，这一步类似于java的全局异常抓取处理
func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	//传入的参数string类型转换为int类型，转换不了，返回异常
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return
	}
	//调用斐波那契列方法，异常不为空，返回异常
	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(list)

}

//测试
func TestGetFibonacci(t *testing.T) {
	//GetFibonacci2("test")//运行结果，类型转换失败异常  Error strconv.Atoi: parsing "test": invalid syntax
	//GetFibonacci2("1")//运行结果，参数校验异常 Error 参数小于 2
	//GetFibonacci2("101") //运行结果，参数校验异常 Error 参数大于 100
	GetFibonacci2("10") //运行结果，运行正常，实现斐波那契列 [1 1 2 3 5 8 13 21 34 55]
}
