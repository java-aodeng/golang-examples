package go_30

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

/*
对象池

go语言对象池 可以利用Buffered Channels来实现

*/

//定义一个对象
type ReusableObj struct {
}

//用于缓冲可重用对象
type ObjPool struct {
	bufChan chan *ReusableObj
}

func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	//创建的Buffered Channels 指定容量numOfObj
	objPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout): //超时控制
		return nil, errors.New("超时")
	}
}

func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")

	}
}

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
	fmt.Println("Done")
}

/*运行结果
=== RUN   TestObjPool
*go_30.ReusableObj
*go_30.ReusableObj
*go_30.ReusableObj
*go_30.ReusableObj
*go_30.ReusableObj
*go_30.ReusableObj
*go_30.ReusableObj
*go_30.ReusableObj
*go_30.ReusableObj
*go_30.ReusableObj
*go_30.ReusableObj
Done
--- PASS: TestObjPool (0.00s)
PASS


Process finished with the exit code 0

*/
