package go_38

import (
	"fmt"
	"reflect"
	"testing"
)

/*
pipe-filter模式
- 非常适合与数据处理及数据分析系统
- Filter封装数据处理的能力
- 松耦合：Filter只跟数据（格式）耦合
- Pipe用于连接Filter传递数据或者在异步处理过程中缓冲数据流进程内同步调用时，pipe演变为数据在方法调用传递

*/

func TestStringSplit(t *testing.T) {
	//定义分隔符号
	sf := NewSplitFilter(",")
	//定义参数
	resp, err := sf.Process("1,2,3")
	if err != nil {
		t.Fatal(err)
	}
	parts, ok := resp.([]string)
	if !ok {
		t.Fatalf("Repsonse type is %T, but the expected type is string", parts)
	}
	if !reflect.DeepEqual(parts, []string{"1", "2", "3"}) {
		t.Errorf("Expected value is {\"1\",\"2\",\"3\"}, but actual is %v", parts)
	}
	t.Error(parts)
}

/*参数是string类型，处理了
=== RUN   TestStringSplit
    pipe_filter_test.go:32: [1 2 3]
--- FAIL: TestStringSplit (0.00s)

FAIL
*/

func TestWrongInput(t *testing.T) {
	sf := NewSplitFilter(",")
	a, err := sf.Process(123)
	if err == nil {
		t.Fatal("An error is expected.")
	}
	fmt.Println(a)
}

/*运行结果，参数不是string类型，不会处理
=== RUN   TestWrongInput
<nil>
--- PASS: TestWrongInput (0.00s)
PASS
*/
