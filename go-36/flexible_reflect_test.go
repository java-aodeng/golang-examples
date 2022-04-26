package go_36

import (
	"errors"
	"reflect"
	"testing"
)

/*
万能程序（原理都是利用go反射的特性，就和java差球不多，动态编程而已，理解了就是有手就行，看不懂这边建议直接放弃）

DeepEqual
- 比较切片和map

万能赋值
- 给两个不同的对象，但是相同属性的字段同时赋值

两个万能程序示例如下：
*/

/*
DeepEqual =====================万能程序1===========================
- 比较切片和map
*/
func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}
	//t.Log(a == b) //运行这行会报错，go语言不能直接比较 invalid operation: a == b (map can only be compared to nil)
	t.Log(reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}
	t.Log("s1 == s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3?", reflect.DeepEqual(s1, s3))
}

/* 运行结果：可以看到我们可以使用DeepEqual比较两个数组是否相等
=== RUN   TestDeepEqual
flexible_reflect_test.go:21: true
flexible_reflect_test.go:26: s1 == s2? true
flexible_reflect_test.go:27: s1 == s3? false
--- PASS: TestDeepEqual (0.00s)
PASS
*/

/*
=====================万能程序2===========================
万能赋值
- 给两个不同的对象，但是相同属性的字段同时赋值。
*/

//下面两个对象：相同属性为：Name，Age
type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

//给不同对象，相同属性赋值
//st参数为传入的类型，settings为传入的值
func fillBySettings(st interface{}, settings map[string]interface{}) error {

	// func (v Value) Elem() Value
	// Elem returns the value that the interface v contains or that the pointer v points to.
	// It panics if v's Kind is not Interface or Ptr.
	// It returns the zero Value if v is nil.

	if reflect.TypeOf(st).Kind() != reflect.Ptr { //判断参数是不是一个指针
		return errors.New("the first param should be a pointer to the struct type.")
	}
	// Elem() 获取指针指向的值
	if reflect.TypeOf(st).Elem().Kind() != reflect.Struct { //判断参数是不是一个结构体
		return errors.New("the first param should be a pointer to the struct type.")
	}

	if settings == nil {
		return errors.New("settings is nil.")
	}

	var (
		field reflect.StructField
		ok    bool
	)

	//这里其实就是循环遍历
	for k, v := range settings {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		//如果传入的对象的属性 和值的key相同就赋值
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}

	}
	return nil
}

func TestFillNameAndAge(t *testing.T) {
	//定义一个数组
	settings := map[string]interface{}{"Name": "Mike", "Age": 30}
	e := Employee{}
	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}

/*
运行结果：两个对象，相同的属性赋值相同了，这样有的场景下，能节省很多代码
=== RUN   TestFillNameAndAge
    flexible_reflect_test.go:114: { Mike 30}
    flexible_reflect_test.go:119: { Mike 30}
--- PASS: TestFillNameAndAge (0.00s)
PASS

*/
