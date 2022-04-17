package go_27

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

/*
其实就是java里面的单例模式，一样理解即可
go语言里面有个once方法只会执行一次，使用这个实现线程安全的懒汉单例模式
*/

type Singleton struct {
	data string
}

var sigleInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {

	//这个方法只会执行一次，不需要像java一样非空判断,第一次是肯定会创建对象的，且只会执行一次
	once.Do(func() {
		fmt.Println("Create Obj")
		sigleInstance = new(Singleton)
	})
	return sigleInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("对象引用地址：%X\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}

/*
运行结果：可以看到对象只创建了一次，后面都是获取的同一个对象引用，与java单例模式一样理解即可
=== RUN   TestGetSingletonObj
Create Obj
对象引用地址：C000030010
对象引用地址：C000030010
对象引用地址：C000030010
对象引用地址：C000030010
对象引用地址：C000030010
对象引用地址：C000030010
对象引用地址：C000030010
对象引用地址：C000030010
对象引用地址：C000030010
对象引用地址：C000030010
--- PASS: TestGetSingletonObj (0.00s)
PASS
*/
