package go_34

/*
Behavior Driven Development

BDD in Go
开源地址：github.com/smartystreets/goconvey

安装
$ go get github.com/smartystreets/goconvey

启动 WEB UI
$ $GOPATH/bin/goconvey

这里运行需要go的1.14版本，不然会出现 errors.Is的问题
*/
import (
	. "github.com/smartystreets/goconvey"
	"testing"
)

func TestSpec(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given 2 even numbers", t, func() {
		a := 3
		b := 4

		Convey("When add the two numbers", func() {
			c := a + b

			Convey("Then the result is still even", func() {
				So(c%2, ShouldEqual, 0)
			})
		})
	})
}