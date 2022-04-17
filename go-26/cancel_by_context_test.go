package go_25

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/*
Context
关联任务的取消，就是用来当前线程被取消，相关联的子线程也会被取消的场景
- 根Context：通过context.Background()创建
- 子Context：context.WithCancel(parentContext)创建
	ctx,cancel=context.WithCancel(context.Background())
- 当前Context被取消时，基于他的子context都会被取消
- 接受取消通知<-ctx.Done()

*/
func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled(ctx) {
					fmt.Println("isCancelled")
					break
				}
				fmt.Println("isSleep")
				time.Sleep(time.Millisecond * 1)
			}
			fmt.Println(i, "Cancelled")
		}(i, ctx)
	}
	//关闭渠道 关闭之后上面多线程里面会自动关闭循环
	cancel()

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
