package go_28

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

/*
仅需任意任务完成

举例：我们搜索一个问题，同时触发google，百度，必应三个程序去搜索，只要有一个程序搜索到结果，就可以返回结果了

*/

//一个简单的函数，打印传入的id
func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("当前返回id为%d", id)
}

//创建一个协程，返回第一个传入的数据id
func FirstResponse() string {
	numOfRunner := 10
	//这里创建一个Buffered Channels，防止协程阻塞，很重要
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestFirsResponse(t *testing.T) {
	fmt.Println("Before", runtime.NumGoroutine())
	fmt.Println(FirstResponse())
	time.Sleep(time.Second * 1)
	fmt.Println("After", runtime.NumGoroutine())
}

/*运行结果 返回id会每次不相同，这是因为go的协程调用机制的顺序不同
=== RUN   TestFirsResponse
Before 2
当前返回id为2
After 2
--- PASS: TestFirsResponse (1.02s)
PASS
*/
