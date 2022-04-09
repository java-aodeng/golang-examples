package go_22

import (
	"fmt"
	"testing"
	"time"
)

/*
CSP并发机制
当成java的异步操作理解就行了,而CSP模式则是通过Channel进行通讯的，更松耦合一些

Channel（管道，像理解消息队列的管道一样理解就行了）
- Channel是Go中的一个核心类型，你可以把它看成一个管道，通过它并发核心单元就可以发送或者接收数据进行通讯
- 容量(capacity)代表Channel容纳的最多的元素的数量，代表Channel的缓存的大小。
- 如果没有设置容量，或者容量设置为0, 说明Channel没有缓存，只有sender和receiver都准备好了后它们的通讯

##如图：普通Channel 与 设置容量了Buffered Channels的区别
图左是普通Channel  图右是设置容量了Buffered Channels
图片路径：https://github.com/java-aodeng/golang-examples/blob/master/go-22/channel.png

*/

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

//运行结果 可以看到结果是串行的，而且时间也等于两个函数总数相加 0.15s
//=== RUN   TestService
//Done
//working on something else
//Task is done.
//--- PASS: TestService (0.15s)
//PASS

func AsyncService() chan string {
	//使用make初始化普通Channel,不指定容量默认就是0 就是普通Channel
	retCh := make(chan string)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func AsyncService2() chan string {
	//使用make初始化Buffered Channels, 只要设置了容量，就是Buffered Channels
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

//使用普通Channel
func TestAsynService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
}

//运行结果 可以看到service exited.是最后返回结果的
//=== RUN   TestAsynService
//working on something else
//returned result.
//Task is done.
//Done
//service exited.
//--- PASS: TestAsynService (0.10s)
//PASS

//使用指定容量的Buffered Channels
func TestAsynService2(t *testing.T) {
	retCh := AsyncService2()
	otherTask()
	fmt.Println(<-retCh)
}

//运行结果 可以看到在等待的时候同时执行了，service exited.没有阻塞
//=== RUN   TestAsynService
//working on something else
//returned result.
//service exited.
//Task is done.
//Done
//--- PASS: TestAsynService (0.10s)
//PASS
