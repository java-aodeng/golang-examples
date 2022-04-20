package go_31

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

/*
sync.pool对象缓存

sync.Pool对象获取
- 尝试从私有对象获取
- 私有对象不存在，尝试从当前Processor的共享池获取
- 如果当前Processor共享池是空的，那么就尝试去其他Processor的共享池获取
- 如果所有池都是空的，最后就用用户指定的new函数产生一个新的对象返回

sync.Pool对象的生命周期
- GC会清楚sync.pool缓存的对象
- 对象的缓存有效期为下一次GC之前

sync.Pool总结
- 使用与通过复用，降低复杂对象的创建和GC代价
- 协程安全，会有锁的开销
- 生命周期受GC影响，不适合于做连接池等，需要自己管理生命周期的资源池化
*/

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("创建一个对象.")
			return 100
		},
	}
	v := pool.Get().(int)
	fmt.Println("得到的对象", v)
	//手动添加值
	pool.Put(1)
	//GC会清除sync.pool中缓存的对象
	runtime.GC()
	v1, _ := pool.Get().(int)
	fmt.Println("v1=", v1)
}

/*TestSyncPool运行结果 可以看到GC之前put的1没有输出，被清除了
=== RUN   TestSyncPool
创建一个对象.
得到的对象 100
创建一个对象.
v1= 100
--- PASS: TestSyncPool (0.00s)
PASS
*/

func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("创建一个对象.")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}

/*
运行结果 可以看到10次里面三次都是取的100，7次是创建的10
=== RUN   TestSyncPoolInMultiGroutine
100
创建一个对象.
10
100
创建一个对象.
10
创建一个对象.
创建一个对象.
10
创建一个对象.
创建一个对象.
10
创建一个对象.
10
100
10
10
*/
