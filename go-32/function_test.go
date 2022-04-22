package go_32

import (
	"fmt"
	"testing"
)

/*
内置单元测试框架
go语言创建单元测试文件，就是文件命名加个_test，测试的入口方法Test开头就行了，比java简洁多了

- Fail,Error：该测试失败，该测试继续，其他测试继续执行
- FailNow,Fatal：该测试失败，该测试中止，其他测试继续执行

*/

func TestErrorInCode(t *testing.T) {
	fmt.Println("Start")
	t.Error("Error")
	fmt.Println("End")
}

func TestFailInCode(t *testing.T) {
	fmt.Println("Start")
	t.Fatal("Error")
	fmt.Println("End")
}

/*
运行结果：
=== RUN   TestErrorInCode
Start
End
--- FAIL: TestErrorInCode (0.00s)
    function_test.go:17: Error
=== RUN   TestFailInCode
Start
--- FAIL: TestFailInCode (0.00s)
    function_test.go:23: Error
FAIL

*/
