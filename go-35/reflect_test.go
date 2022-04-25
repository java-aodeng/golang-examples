package go_35

import (
	"fmt"
	"reflect"
	"testing"
)

/*
反射编程

reflect.TypeOf   reflect.ValueOf
- reflect.TypeOf 返回类型(reflect.Type)
- reflect.ValueOf 返回值(reflect.Value)
- 可以从reflect.Value获得类型
- 通过kind来判断类型
*/

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

/*
运行结果：
=== RUN   TestTypeAndValue
    reflect_test.go:21: int64 10
    reflect_test.go:22: int64
--- PASS: TestTypeAndValue (0.00s)
PASS
*/

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown", t)
	}
}

//用反射获得类型
func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(&f)

}

/*
运行结果
=== RUN   TestBasicType
Unknown *float64
--- PASS: TestBasicType (0.00s)
PASS
*/

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}
	//	t.Log(a == b)
	t.Log("a==b?", reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}

	t.Log("s1 == s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3?", reflect.DeepEqual(s1, s3))

	c1 := Customer{"1", "Mike", 40}
	c2 := Customer{"1", "Mike", 40}
	fmt.Println(c1 == c2)
	fmt.Println(reflect.DeepEqual(c1, c2))
}

/*
运行结果：
=== RUN   TestDeepEqual
    reflect_test.go:47: a==b? true
    reflect_test.go:53: s1 == s2? true
    reflect_test.go:54: s1 == s3? false
true
true
--- PASS: TestDeepEqual (0.00s)
PASS
*/

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	//按名字获取成员
	t.Logf("Name: value(%[1]v), Type(%[1]T) ", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name' field.")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").
		Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Updated Age:", e)
}

/*
运行结果
=== RUN   TestInvokeByName
    reflect_test.go:81: Name: value(Mike), Type(reflect.Value)
    reflect_test.go:85: Tag:format normal
    reflect_test.go:89: Updated Age: &{1 Mike 1}
--- PASS: TestInvokeByName (0.00s)
PASS
*/
