package go_23

import (
	"fmt"
	"testing"
	"time"
)

/*
多路由选择和超时
*/

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func AsyncService() chan string {
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

func TestSelect(t *testing.T) {
	//select 多渠道选择，语法与switch类似
	select {
	case ret := <-AsyncService():
		t.Log(ret)
		//如果超时执行这个
	case <-time.After(time.Millisecond * 100):
		t.Errorf("time out")
		//default:
		//	t.Log("执行完成")
	}
}

//运行结果
//=== RUN   TestSelect
//--- FAIL: TestSelect (0.10s)
//select_test.go:35: time out
//
//FAIL
