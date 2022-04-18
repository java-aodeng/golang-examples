package go_29

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

/*
所有任务完成

与上一章刚好相反，要程序所有协程执行完了，才返回结果
*/

//一个简单的函数，打印传入的id
func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("当前返回id为%d", id)
}

//所有任务完成才返回
func AllResponse() string {
	numOfRunner := 10
	//这里创建一个Buffered Channels，防止协程阻塞，很重要
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}

	finalRet := ""
	for j := 0; j < numOfRunner; j++ {
		finalRet += <-ch + "\n"
	}

	return finalRet
}

func TestFirsResponse(t *testing.T) {
	fmt.Println("Before", runtime.NumGoroutine())
	fmt.Println(AllResponse())
	time.Sleep(time.Second * 1)
	fmt.Println("After", runtime.NumGoroutine())
}

/*
运行结果
=== RUN   TestFirsResponse
Before 2
当前返回id为9
当前返回id为7
当前返回id为1
当前返回id为8
当前返回id为2
当前返回id为5
当前返回id为4
当前返回id为6
当前返回id为0
当前返回id为3

After 2
--- PASS: TestFirsResponse (1.01s)
PASS
*/
