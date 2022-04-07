package go_21

import (
	"sync"
	"testing"
	"time"
)

/*
共享内存并发机制 就类似于java里面的锁

都是些很简单的东西，直接看代码

*/

//不加锁的情况
func TestCounter(t *testing.T) {
	count := 0
	for i := 0; i < 5000; i++ {
		go func() {
			count++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("count=%d", count)
}

//运行结果 结果不是我们预想中的5000，发生了写入的问题，不是线程安全的
//=== RUN   TestCounter
//--- PASS: TestCounter (1.02s)
//share_mem_test.go:22: count=4936
//PASS

//加锁的情况
func TestCounterThreadSafe(t *testing.T) {
	var mut = sync.Mutex{}
	count := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			count++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("count=%d", count)
}

//运行结果 加锁之后结果是预想中的
//=== RUN   TestCounterThreadSafe
//--- PASS: TestCounterThreadSafe (1.02s)
//share_mem_test.go:47: count=5000
//PASS

//线程等待
func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			count++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("count=%d", count)
}

//运行结果
//=== RUN   TestCounterWaitGroup
//--- PASS: TestCounterWaitGroup (0.03s)
//share_mem_test.go:73: count=5000
//PASS
