package go_20

import (
	"fmt"
	"testing"
	"time"
)

/*

协程机制
就是一种更轻量级的线程
其实你可以理解为java的多线程编程，只不过性能处理更加好

java Thead 对比 go Groutine
## 创建时默认的stack（栈）的大小
1.jdk5以后java thread stack默认为1M
2.Groutine的stack初始大小为2k（这也就是为什么go语言比java更适合并发编程了）

## 和KSE（kernel space entity 系统线程）的对应关系
1.java thread是1:1
2.Groutine是M:N（多对多，所以go天生就是用来处理并发的）

*/

//测试
func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		// 创建协程，如下写一个go就行了
		go func(i int) {
			fmt.Println(i)
		}(i) //这里i是值传递，每次都是复制一个，所以不会存在共享变量，需要像java加锁一样的问题，所以go天生并发
	}
	time.Sleep(time.Millisecond * 50)
}

/*
运行结果 结果顺序不一样，和java多线程一样理解就行了
=== RUN   TestGroutine
2
0
5
3
4
1
7
6
8
9
--- PASS: TestGroutine (0.05s)
PASS
*/
