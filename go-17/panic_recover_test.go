package go_17

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"
)

//panic
//用于主动抛出错误,类似于java的抛出异常
//panic用于不可以恢复的错误
//panic退出前回执行defer指定的内容

//panic 和os.Exit的区别
//os.Exit退出时不会调用defer指定的函数
//os.Exit退出时不输出当前调用栈信息

//recover
//用于捕获panic抛出的错误，类似于java的异常捕获try{}catch

//使用Exit退出程序
func TestExit(t *testing.T) {
	defer func() {
		fmt.Print("我是退出前回执的内容")
	}()
	fmt.Print("Start--》")
	os.Exit(-1)
}

//运行结果 ：可以看到退出后没有执行defer里面的内容
//=== RUN   TestExit
//Start--》

//使用panic抛出异常
func TestPainc(t *testing.T) {
	defer func() {
		fmt.Print("我是退出前回执的内容")
	}()
	fmt.Print("Start--》")
	panic(errors.New("异常退出"))
}

//运行结果 ：可以看到退出后执行defer里面的内容，并且输出当前调用栈信息
//=== RUN   TestPainc
//Start--》我是退出前回执的内容--- FAIL: TestPainc (0.00s)
//panic: 异常退出 [recovered]
//panic: 异常退出
//goroutine 19 [running]:
//testing.tRunner.func1(0xc0000a0100)

//使用recover捕获panic抛出的异常
func TestRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			//这里可以做"错误恢复"
			//常见的的"错误恢复"，记录到打印日志里面,但是这样的修复方式是非常危险的，很容易出先僵尸进程，”Let it Crash“往往是我们恢复不确定性错误的最好方法
			log.Println("捕获到异常，手动处理异常")
		}
	}()
	fmt.Print("Start")
	panic(errors.New("异常退出"))
}

//运行结果 大家可以看到这里输出内容不是"异常退出"也没有打印栈信息了，因为我们手动处理了异常，但是这样做是对程序很危险的一种方式，懒点直接让程序崩溃反而更好，不然服务器炸了，不要让运维有甩锅给开发的机会
//=== RUN   TestRecover
//Start2022/04/02 19:21:36 捕获到异常，手动处理异常信息
//--- PASS: TestRecover (0.03s)
//PASS
