package go_24

import (
	"fmt"
	"sync"
	"testing"
)

/*
channel的关闭
- 向关闭的channel发送数据，会导致panic
- v,ok<-ch;ok为bool值，true表示正常接受，false表示通道关闭
- 所有的channel接收者都会再channel关闭时，立刻从阻塞等待中返回且上述ok值为false，这个广播机制常被利用，进行向多个订阅者同时发送信息。如：退出信号。

*/

//数据提供者 向chan里面传0-9的数据
func dataProduct(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		//执行完了，关闭chan
		close(ch)
		//ch<-11 向关闭的chan发消息会抛异常
		//线程安全，不是本节重点，看不懂忽略
		wg.Done()
	}()
}

//数据消费者 取出chan里面0-9的数据 输出
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

//测试
func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	//这里创建的是普通chan，不会存关闭chan有缓存的问题
	ch := make(chan int)

	//调用数据提供者 设置chan
	wg.Add(1)
	dataProduct(ch, &wg)

	//调用数据消费者  消费赋值了的chan
	wg.Add(1)
	dataReceiver(ch, &wg)

	//调用数据消费者  消费赋值了的chan
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}

/*
运行结果 可以看到输出顺序不一样，因为我们调用了两个消费者 还是很简单的东西
=== RUN   TestCloseChannel
0
1
2
3
5
6
7
8
9
4
--- PASS: TestCloseChannel (0.00s)
PASS


*/
