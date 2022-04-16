package go_25

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cacelCh chan struct{}) {
			for {
				if isCancelled(cancelChan) {
					fmt.Println("isCancelled")
					break
				}
				fmt.Println("isSleep")
				time.Sleep(time.Millisecond * 1)
			}
			fmt.Println(i, "Cancelled")
		}(i, cancelChan)
	}
	//关闭渠道 关闭之后上面多线程里面会自动关闭循环
	//cancel_2(cancelChan)

	time.Sleep(time.Second * 10)
}

/*
运行结果 执行cancel_2方法后关闭了渠道 多线程for循环里面就立马关闭了循环，注释掉cancel_2再执行，会for循环打印isSleep直到10秒后退出
=== RUN   TestCancel
isCancelled
3 Cancelled
isCancelled
4 Cancelled
isCancelled
1 Cancelled
isCancelled
2 Cancelled
isCancelled
0 Cancelled
--- PASS: TestCancel (1.00s)
PASS


Process finished with the exit code 0

*/
